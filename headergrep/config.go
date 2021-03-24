package headergrep

import "net/http"

type Config struct {
	Expected    http.Header
	UnExpected  http.Header
	IgnoreHTTPS bool
	Method      string
}
