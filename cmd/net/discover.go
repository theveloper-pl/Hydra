package net

import (
	"fmt"
	"github.com/spf13/cobra"
)

var discoverCmd = &cobra.Command{
	Use:   "discover",
	Short: "Discover specified host to test its reachability on an IP network",
	Long: `Valuable tool for obtaining information about domain names and IP addresses`,
	Run: func(cmd *cobra.Command, args []string) {
		do_discover(host_ip_domain)
	},
}

func init() {
	discoverCmd.Flags().StringVarP(&host_ip_domain, "host", "", "", "Host's domain or ip to ping")

	if err := discoverCmd.MarkFlagRequired("host"); err != nil{
		fmt.Println(err)
	}

	
	NetCmd.AddCommand(discoverCmd)
}
