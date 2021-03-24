package headergrep

import "net/http"

type Config struct {
	Expected    http.Header
	UnExpected  http.Header
	IgnoreHTTPS bool
	Method      string
}

func NewConfig() Config {
	return Config{
		Expected:   make(http.Header),
		UnExpected: make(http.Header),
	}
}
