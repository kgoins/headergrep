package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/kgoins/headergrep/headergrep"
	"github.com/kgoins/snakecharmer/snakecharmer"
	homedir "github.com/mitchellh/go-homedir"
)

var rootCmd = &cobra.Command{
	Use:   "headergrep url",
	Short: "A small tool for examining headers from a web request",
	Args:  cobra.MinimumNArgs(1),
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		sc := snakecharmer.NewSnakeCharmer("headergrep", ".headergrep")
		sc.InitConfig(cmd, filepath.Join(home, ".headergrep"))
		return
	},
	Run: func(cmd *cobra.Command, args []string) {
		config, err := BuildConfigFromCmd(cmd)
		if err != nil {
			log.Fatalln("Unable to build config: " + err.Error())
		}

		url, err := url.Parse(args[0])
		if err != nil {
			log.Fatalln("Unable to parse url: " + err.Error())
		}

		hg := headergrep.NewHGrep(config)
		respHeaders, err := hg.GetHeaders(url)
		if err != nil {
			log.Fatalln("Unable to get headers: " + err.Error())
		}

		err = PrintRespHeaders(respHeaders)
		if err != nil {
			log.Fatalln("Unable to print headers: " + err.Error())
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP(
		"config",
		"c",
		"",
		"config file (default is $HOME/.headergrep.toml)",
	)

	rootCmd.PersistentFlags().BoolP(
		"ignorehttps",
		"k",
		false,
		"ignore https cert issues",
	)

	rootCmd.PersistentFlags().StringP(
		"method",
		"X",
		"GET",
		"specify http method to use",
	)

	rootCmd.PersistentFlags().StringArrayP(
		"expected",
		"e",
		[]string{},
		"headers that are expected to be present",
	)

	rootCmd.PersistentFlags().StringArrayP(
		"unexpected",
		"u",
		[]string{},
		"headers that are expected to be absent",
	)
}
