package shodan

import "fmt"

// APIError is raised whenever a non-200 status code was returned by the Shodan API.
type APIError struct {
	Value string
}

func (e *APIError) Error() string {
	if e == nil {
		return ""
	}
	return e.Value
}

func newAPIError(value string) error {
	return &APIError{Value: value}
}

func newAPIErrorf(format string, args ...any) error {
	return &APIError{Value: fmt.Sprintf(format, args...)}
}

// APITimeout is a timed-out API request.
type APITimeout struct {
	APIError
}
