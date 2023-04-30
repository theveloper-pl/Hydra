package net

import (
	"fmt"
	"github.com/spf13/cobra"
)

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Ping specified host to test its reachability on an IP network",
	Long: `Network utility that tests the connectivity between two devices by sending ICMP echo request packets and measuring response times`,
	Run: func(cmd *cobra.Command, args []string) {
		do_ping(host_ip_domain, pings_number)
	},
}

func init() {
	pingCmd.Flags().StringVarP(&host_ip_domain, "host", "", "", "Host's domain or ip to ping")
	pingCmd.Flags().IntVarP(&pings_number, "count", "", 1, "Amount of pings that will be sent")

	if err := pingCmd.MarkFlagRequired("host"); err != nil{
		fmt.Println(err)
	}

	
	NetCmd.AddCommand(pingCmd)
}
