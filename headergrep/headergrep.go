package headergrep

import (
	"crypto/tls"
	"net/http"
	"net/url"
)

// headerPlaceholder represents an "empty" header value
// since http.Header.Get() returns "" if a header doesn't exist
const headerPlaceholder string = " "

type HeaderGrepper interface {
	GetHeaders(url *url.URL) (RespHeaders, error)
}

type HGrep struct {
	Config
}

// Verify that HGrep implements HeaderGrepper.
var _ HeaderGrepper = HGrep{}
var _ HeaderGrepper = (*HGrep)(nil)

func NewHGrep(config Config) HGrep {
	return HGrep{Config: config}
}

func (h HGrep) GetHeaders(url *url.URL) (respHeaders RespHeaders, err error) {
	req, err := http.NewRequest(h.Method, url.String(), nil)
	if err != nil {
		return
	}

	client := http.Client{}
	if h.Config.IgnoreHTTPS {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}

		client.Transport = tr
	}

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	return h.BuildRespHeaders(resp), nil
}

func (h HGrep) isExpected(key string) bool {
	return h.Config.Expected.Get(key) == headerPlaceholder
}

func (h HGrep) isUnExpected(key string) bool {
	return h.Config.UnExpected.Get(key) == headerPlaceholder
}

func (h HGrep) BuildRespHeaders(resp *http.Response) RespHeaders {
	respHeaders := NewResponseHeaders()

	for name, val := range resp.Header {
		if h.isExpected(name) {
			respHeaders.AddExpected(name, val)
		} else if h.isUnExpected(name) {
			respHeaders.AddUnExpected(name, val)
		} else {
			respHeaders.AddOther(name, val)
		}
	}

	return respHeaders
}
