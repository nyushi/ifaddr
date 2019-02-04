package cmd

import (
	"log"

	"github.com/nyushi/ifaddr"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addrListCmd)
}

var addrListCmd = &cobra.Command{
	Use:   "addrlist",
	Short: "List addresses",
	Run: func(cmd *cobra.Command, args []string) {
		//addrs, err := ifaddr.ListAddresses(flagV4, flagV6, flagLoopback, flagLinklocal, flagMulticast)
		addrs, err := ifaddr.ListAddresses(ifaddr.IPTypeQuery{
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
