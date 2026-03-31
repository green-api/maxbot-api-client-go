package maxerr

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrBadRequest         = errors.New("bad request")         /* 400 */
	ErrUnauthorized       = errors.New("unauthorized")        /* 401 */
	ErrNotFound           = errors.New("not found")           /* 404 */
	ErrMethodNotAllowed   = errors.New("method not allowed")  /* 405 */
	ErrTooManyRequests    = errors.New("too many requests")   /* 429 */
	ErrServiceUnavailable = errors.New("service unavailable") /* 503 */
)

func HandleErrorResponse(resp *http.Response, body []byte) error {
	var baseErr error

	switch resp.StatusCode {
	case http.StatusBadRequest:
		baseErr = ErrBadRequest
	case http.StatusUnauthorized:
		baseErr = ErrUnauthorized
	case http.StatusNotFound:
		baseErr = ErrNotFound
	case http.StatusMethodNotAllowed:
		baseErr = ErrMethodNotAllowed
	case http.StatusTooManyRequests:
		baseErr = ErrTooManyRequests
	case http.StatusServiceUnavailable:
		baseErr = ErrServiceUnavailable
	default:
		baseErr = fmt.Errorf("unexpected API error")
	}

	return fmt.Errorf("%w: status code %d, response: %s", baseErr, resp.StatusCode, string(body))
}
