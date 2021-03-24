package main

import (
	"bufio"
	"net/http"
	"net/textproto"
	"strings"

	"github.com/kgoins/headergrep/headergrep"
	"github.com/spf13/cobra"
)

func headersFromStrings(strs []string) (http.Header, error) {
	if len(strs) == 0 {
		return make(http.Header), nil
	}

	headerStr := strings.Join(strs, "\r\n")
	strReader := bufio.NewReader(strings.NewReader(headerStr))
	headerReader := textproto.NewReader(strReader)

	header, err := headerReader.ReadMIMEHeader()
	if err != nil {
		return nil, err
	}

	return http.Header(header), nil
}

func BuildConfigFromCmd(cmd *cobra.Command) (headergrep.Config, error) {
	config := headergrep.NewConfig()
	config.IgnoreHTTPS, _ = cmd.Flags().GetBool("ignorehttps")
	config.Method, _ = cmd.Flags().GetString("method")

	expectedRaw, _ := cmd.Flags().GetStringArray("expected")
	unexpectedRaw, _ := cmd.Flags().GetStringArray("unexpected")

	var err error
	config.Expected, err = headersFromStrings(expectedRaw)
	if err != nil {
		return config, err
	}

	config.UnExpected, err = headersFromStrings(unexpectedRaw)
	if err != nil {
		return config, err
	}

	return config, nil
}
