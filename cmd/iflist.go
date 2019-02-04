package cmd

import (
	"log"

	"github.com/nyushi/ifaddr"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(iflistCmd)
}

var iflistCmd = &cobra.Command{
	Use:   "iflist",
	Short: "List interfaces",
	Run: func(cmd *cobra.Command, args []string) {
		names, err := ifaddr.ListInterfaces()
		if err != nil {
			log.Fatalf("failed to get interface names: %s", err)
		}
		PrintOutput(names, flagJSON)
	},
}
