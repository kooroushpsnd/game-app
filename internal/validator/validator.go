package validator

import (
	"goProject/internal/pkg/errmsg"
	"goProject/internal/pkg/richerror"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type AppValidator struct {
	v *validator.Validate
}

func New() *AppValidator {
	v := validator.New()

	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("json")
		if name == "" {
			return fld.Name
		}
		name = strings.Split(name, ",")[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return &AppValidator{v: v}
}

func (a *AppValidator) Validate(i any) error {
	const op = "validator"
	if err := a.v.Struct(i);err != nil {
		return richerror.New(op).
			WithErr(err).
			WithKind(richerror.KindInvalid).
			WithMessage(errmsg.ErrorMsg_InvalidInput).
			WithMeta(map[string]interface{}{
				"request": i,
			})
	}
	return nil
}
