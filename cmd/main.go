package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/theblitlabs/parity-server/cmd/cli"
	"github.com/theblitlabs/parity-server/pkg/logger"
)

var logMode string

var rootCmd = &cobra.Command{
	Use:   "parity-server",
	Short: "Parity Protocol Server",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		switch logMode {
		case "debug", "pretty", "info", "prod", "test":
			logger.InitWithMode(logger.LogMode(logMode))
		default:
			logger.InitWithMode(logger.LogModePretty)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		cli.RunServer()
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func ExecuteServer() error {
	serverCmd.Run(serverCmd, []string{})
	return nil
}

func init() {
	rootCmd.PersistentFlags().StringVar(&logMode, "log", "pretty", "Log mode: debug, pretty, info, prod, test")
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the parity server",
	Run: func(cmd *cobra.Command, args []string) {
		cli.RunServer()
	},
}
