package config

import "sync"

type httpError struct {
	Code        int
	ServiceCode int
	Detail      string
}

type HTTPErrors struct {
	Internal httpError
}

func setHTTPErrors(errors *HTTPErrors) {
	errors.Internal = httpError{
		Code:        500,
		ServiceCode: 0,
		Detail:      "Internal Error",
	}
}

func GetHTTPErrors() *HTTPErrors {
	var errors HTTPErrors
	var once sync.Once

	once.Do(func() {
		setHTTPErrors(&errors)
	})

	return &errors
}
