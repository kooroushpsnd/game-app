package postgresemailcode

import (
	"goProject/internal/entity"
	"goProject/internal/repository/postgres"
)

func scanEmailCode(scanner postgres.Scanner) (entity.EmailCode, error) {
	var emailCode entity.EmailCode

	err := scanner.Scan(
		&emailCode.ID ,
		&emailCode.Email ,
		&emailCode.HashCode ,
		&emailCode.Status ,
		&emailCode.Attempts ,
		&emailCode.ExpirationDate ,
		&emailCode.UserID ,
		&emailCode.CreatedAt,
	)

	return emailCode, err
}