/*
Copyright 2019 Adevinta
*/

package cmd

import (
	"fmt"
	"os"

	cli "github.com/adevinta/vulcan-core-cli"

	"github.com/spf13/cobra"
)

var (
	scheme, host string
	verbose      bool
	c            *cli.CLI
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "vulcan-core-cli",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&scheme, "scheme", "s", "https", "scheme of the vulcan-scan-engine endpoint")
	RootCmd.PersistentFlags().StringVarP(&host, "Host", "H", "127.0.0.1", "vulcan-scan-engine endpoint")
	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "prints verbose information during command execution")
}
