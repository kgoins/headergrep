package headergrep_test

import (
	"net/http"
	"testing"

	"github.com/kgoins/headergrep/headergrep"
	"github.com/stretchr/testify/require"
)

func getTestConfig() headergrep.Config {
	config := headergrep.NewConfig()
	config.SetExpected([]string{"Content-Type"})
	config.SetUnExpected([]string{"Server"})

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

	rq.Len(respHeaders.Expected, 1)
	rq.Len(respHeaders.UnExpected, 1)
	rq.Len(respHeaders.Other, 1)
}
