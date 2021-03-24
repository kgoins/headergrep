package main

import (
	"encoding/json"
	"fmt"

	"github.com/kgoins/headergrep/headergrep"
)

func PrintRespHeaders(respHeaders headergrep.RespHeaders) error {
	jsonBytes, err := json.Marshal(respHeaders)
	if err != nil {
		return err
	}

	fmt.Println(string(jsonBytes))
	return nil
}
