package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ifaddr",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		println(1)
	},
}

var (
	flagJSON      bool
	flagV4        bool
	flagV6        bool
	flagLoopback  bool
	flagLinklocal bool
	flagMulticast bool
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&flagJSON, "json", "j", false, "")
	rootCmd.PersistentFlags().BoolVarP(&flagV4, "v4", "4", true, "")
	rootCmd.PersistentFlags().BoolVarP(&flagV6, "v6", "6", false, "")
	rootCmd.PersistentFlags().BoolVarP(&flagLoopback, "loopback", "", false, "")
	rootCmd.PersistentFlags().BoolVarP(&flagLinklocal, "linklocal", "", false, "")
	rootCmd.PersistentFlags().BoolVarP(&flagMulticast, "multicast", "", false, "")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
