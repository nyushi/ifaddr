package cmd

import (
	"log"

	"github.com/nyushi/ifaddr"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get addresses",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("interface name required.")
		}
		addrs, err := ifaddr.GetAddress(args[0], ifaddr.IPTypeQuery{
			IsLoopback:  flagLoopback,
			IsMulticast: flagMulticast,
			IsLinklocal: flagLinklocal,
			IsIPv4:      flagV4,
			IsIPv6:      flagV6,
		})
		if err != nil {
			log.Fatalf("failed to get interface names: %s", err)
		}
		PrintOutput(addrs, flagJSON)
	},
}
