package headergrep

import (
	"crypto/tls"
	"net/http"
	"net/url"
)

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

func hasHeader(key string, header http.Header) bool {
	return header.Get(key) != ""
}

func (h HGrep) BuildRespHeaders(resp *http.Response) RespHeaders {
	respHeaders := NewResponseHeaders()

	for name, val := range resp.Header {
		if hasHeader(name, h.Config.Expected) {
			respHeaders.AddExpected(name, val)
		} else if hasHeader(name, h.Config.UnExpected) {
			respHeaders.AddUnExpected(name, val)
		} else {
			respHeaders.AddOther(name, val)
		}
	}

	return respHeaders
}
