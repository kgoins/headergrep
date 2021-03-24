package headergrep

import "net/http"

type RespHeaders struct {
	Expected   http.Header
	UnExpected http.Header
	Other      http.Header
}

func NewResponseHeaders() RespHeaders {
	return RespHeaders{
		Expected:   make(http.Header),
		UnExpected: make(http.Header),
		Other:      make(http.Header),
	}
}

func (rh RespHeaders) AddExpected(key string, vals []string) {
	for _, val := range vals {
		rh.Expected.Add(key, val)
	}
}

func (rh RespHeaders) AddUnExpected(key string, vals []string) {
	for _, val := range vals {
		rh.UnExpected.Add(key, val)
	}
}

func (rh RespHeaders) AddOther(key string, vals []string) {
	for _, val := range vals {
		rh.Other.Add(key, val)
	}
}
