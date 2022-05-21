/*
Copyright © 2022 Spotlite

*/
package harok

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "harok",
	Short: "CLI for managing your infrastructure",
	Long:  `harok is a SRE tool created to automate and manage your infrastructure`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}
