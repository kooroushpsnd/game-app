package httpmsg

import (
	"goProject/internal/pkg/errmsg"
	"goProject/internal/pkg/richerror"
	"net/http"
)

func Error(err error) (msg string, code int) {
	switch err.(type) {
	case richerror.RichError:
		re := err.(richerror.RichError)
		msg := re.Message()

		code := mapKindToHTTPStatusCode(re.Kind())

		if code > 500 {
			msg = errmsg.ErrorMsg_SomethingWentWrong
		}

		return msg, code
	default:
		return err.Error(), http.StatusBadRequest
	}
}

var kindToStatus = map[richerror.Kind]int{
	richerror.KindInvalid:         http.StatusUnprocessableEntity,
	richerror.KindUnauthorized:    http.StatusUnauthorized,
	richerror.KindForbidden:       http.StatusForbidden,
	richerror.KindNotFound:        http.StatusNotFound,
	richerror.KindConflict:        http.StatusConflict,
	richerror.KindTooManyRequests: http.StatusTooManyRequests,
	richerror.KindTimeout:         http.StatusGatewayTimeout,
	richerror.KindUnavailable:     http.StatusServiceUnavailable,
	richerror.KindUnexpected:      http.StatusInternalServerError,
}

func mapKindToHTTPStatusCode(kind richerror.Kind) int {
	if code, ok := kindToStatus[kind]; ok {
		return code
	}
	return http.StatusBadRequest
}
