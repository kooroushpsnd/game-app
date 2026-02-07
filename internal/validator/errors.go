package validator

import (
	"errors"
	"fmt"
	"goProject/internal/pkg/errmsg"

	v10 "github.com/go-playground/validator/v10"
)

func FieldErrors(err error) map[string]string {
	out := map[string]string{}

	var ve v10.ValidationErrors
	if !errors.As(err, &ve) {
		out["_error"] = err.Error()
		return out
	}

	for _, fe := range ve {
		field := fe.Field()
		switch fe.Tag() {
		case "required":
			out[field] = errmsg.ErrorMsg_Required
		case "email":
			out[field] = errmsg.ErrorMsg_Email
		case "min":
			out[field] = fmt.Sprintf(errmsg.ErrorMsg_Min, fe.Param())
		case "max":
			out[field] = fmt.Sprintf(errmsg.ErrorMsg_Max, fe.Param())
		default:
			out[field] = errmsg.ErrorMsg_IsInvalid
		}
	}

	return out
}
