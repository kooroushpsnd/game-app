package mailer

import (
	"context"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"fmt"
	"math/big"
	"net/smtp"
	"os"
)

type MailerAdapter struct{}

func NewSMTPEmailAdapter() *MailerAdapter {
	return &MailerAdapter{}
}

func (adaptor *MailerAdapter) generateOTP(ctx context.Context, digits int) (string, error) {
	max := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(digits)), nil) // 10^digits

	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%0*d", digits, n.Int64()), nil
}

func (adaptor *MailerAdapter) HashOTP(email, otp string) (string, error) {
	secret := os.Getenv("OTP_SECRET")
	if secret == "" {
		return "", fmt.Errorf("OTP_SECRET is not set")
	}

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(email))
	mac.Write([]byte(":"))
	mac.Write([]byte(otp))
	return hex.EncodeToString(mac.Sum(nil)), nil
}

func (adaptor *MailerAdapter) VerifyOTP(email, otp, storedHash string) (bool, error) {
	hash, err := adaptor.HashOTP(email, otp)
	if err != nil {
		return false, err
	}
	return subtle.ConstantTimeCompare([]byte(hash), []byte(storedHash)) == 1, nil
}

func (adaptor *MailerAdapter) GenerateAndHashOTP(ctx context.Context, email string, digits int) (string, string, error) {
	otp, err := adaptor.generateOTP(ctx, digits)
	if err != nil {
		return "", "", err
	}
	hash, err := adaptor.HashOTP(email, otp)
	if err != nil {
		return "", "", err
	}
	return otp, hash, nil
}

func (adaptor *MailerAdapter) SendEmail(ctx context.Context ,to, subject, body string) error {
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")
	user := os.Getenv("SMTP_USER")
	pass := os.Getenv("SMTP_PASS")

	addr := fmt.Sprintf("%s:%s", host, port)
	auth := smtp.PlainAuth("", user, pass, host)

	msg := []byte(
		"From: " + user + "\r\n" +
			"To: " + to + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/plain; charset=\"utf-8\"\r\n" +
			"\r\n" +
			body + "\r\n",
	)

	return smtp.SendMail(addr, auth, user, []string{to}, msg)
}
