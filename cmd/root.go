/*
Copyright © 2022 THITIKAN PHAGAMAS <thitikan.phagamas@gmail.com>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "codetest",
	Short: "CodeTest CLI is a command-line tool that display a list of funds",
	Long:  `CodeTest CLI is a command-line tool that displays a list of funds over a specified time period.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}
