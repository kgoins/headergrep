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

func headersFromStrings(headerStrs []string) http.Header {
	headers := make(http.Header)
	if len(headerStrs) == 0 {
		return headers
	}

	for _, name := range headerStrs {
		headers.Add(name, headerPlaceholder)
	}

	return headers
}

func (c *Config) SetExpected(headerNames []string) {
	c.Expected = headersFromStrings(headerNames)
}

func (c *Config) SetUnExpected(headerNames []string) {
	c.UnExpected = headersFromStrings(headerNames)
}
