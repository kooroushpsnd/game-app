package helper

import (
	"fmt"
	"goProject/internal/pkg/richerror"
	"reflect"
	"strings"

	"github.com/labstack/echo/v4"
)

func BindValidateQuery[T any](c echo.Context) (T, error) {
	const op = "validator.query"
	var dst T

	// 1) reject unknown keys
	if err := rejectUnknownKeys(c, &dst); err != nil {
		return dst ,richerror.New(op).
				WithErr(err).
				WithKind(richerror.KindInvalid)
	}

	// 2) bind query params
	if err := (&echo.DefaultBinder{}).BindQueryParams(c, &dst); err != nil {
		return dst ,err
	}

	// 3) validate
	if err := c.Validate(&dst); err != nil {
		return dst, err
	}

	return dst, nil
}

func rejectUnknownKeys(c echo.Context, dst any) error {
	const op = "validator.query"
	
	allowed := map[string]struct{}{}
	collectAllowedKeys(reflect.TypeOf(dst), allowed)

	for key := range c.QueryParams() {
		if _, ok := allowed[key]; !ok {
			return richerror.New(op).
				WithKind(richerror.KindInvalid).
				WithMessage(fmt.Sprintf("unknown query param: %s", key))
		}
	}
	return nil
}

func collectAllowedKeys(t reflect.Type, allowed map[string]struct{}) {
	if t == nil {
		return
	}
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		// include embedded structs (e.g. PaginationDto)
		if f.Anonymous {
			collectAllowedKeys(f.Type, allowed)
		}

		tag := strings.TrimSpace(f.Tag.Get("query"))
		if tag == "" {
			continue
		}
		name := strings.TrimSpace(strings.Split(tag, ",")[0])
		if name == "" || name == "-" {
			continue
		}
		allowed[name] = struct{}{}
	}
}
