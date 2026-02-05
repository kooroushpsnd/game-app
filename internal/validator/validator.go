package validator

import (
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
	return a.v.Struct(i)
}
