package main

import (
	"errors"
	"strings"

	"github.com/kgoins/headergrep/headergrep"
	"github.com/spf13/cobra"
)

func strArrayFromViper(strs []string) ([]string, error) {
	if len(strs) == 0 {
		return []string{}, nil
	}

	if len(strs) != 1 {
		return nil, errors.New("Invalid header format")
	}

	baseStr := strs[0]
	baseStr = strings.Trim(baseStr, "[")
	baseStr = strings.Trim(baseStr, "]")

	headerStrs := strings.Split(baseStr, " ")
	return headerStrs, nil
}

func BuildConfigFromCmd(cmd *cobra.Command) (headergrep.Config, error) {
	config := headergrep.NewConfig()
	config.IgnoreHTTPS, _ = cmd.Flags().GetBool("ignorehttps")
	config.Method, _ = cmd.Flags().GetString("method")

	expectedRaw, _ := cmd.Flags().GetStringArray("expected")
	unexpectedRaw, _ := cmd.Flags().GetStringArray("unexpected")

	var err error
	expected, err := strArrayFromViper(expectedRaw)
	if err != nil {
		return config, err
	}
	config.SetExpected(expected)

	unexpected, err := strArrayFromViper(unexpectedRaw)
	if err != nil {
		return config, err
	}
	config.SetUnExpected(unexpected)

	return config, nil
}
