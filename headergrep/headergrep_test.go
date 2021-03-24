package headergrep_test

import (
	"net/http"
	"testing"

	"github.com/kgoins/headergrep/headergrep"
	"github.com/stretchr/testify/require"
)

func getTestConfig() headergrep.Config {
	config := headergrep.NewConfig()
	config.Expected.Add("Content-Type", "application/json")
	config.UnExpected.Add("Server", "Node")

	return config
}

func TestBuildRespHeaders(t *testing.T) {
	rq := require.New(t)

	resp := &http.Response{}
	resp.Header = make(http.Header)

	resp.Header.Add("Content-Type", "application/json")
	resp.Header.Add("Server", "Node")
	resp.Header.Add("X-Custom", "false")

	config := getTestConfig()
	hg := headergrep.NewHGrep(config)
	respHeaders := hg.BuildRespHeaders(resp)

	rq.Equal(config.Expected, respHeaders.Expected)
	rq.Equal(config.UnExpected, respHeaders.UnExpected)
	rq.Len(respHeaders.Other, 1)
}
