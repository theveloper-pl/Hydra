package net

import (
	"fmt"
	"github.com/spf13/cobra"
)

var only_open bool

var scan_portsCmd = &cobra.Command{
	Use:   "scan-ports",
	Short: "Network utility that probes and identifies open ports on a target system.",
	Long: `Port scanning work by sending various types of packets to a range of ports on a target system and monitoring the responses to determine which ports are open, closed, or filtered.`,
	Run: func(cmd *cobra.Command, args []string) {
		do_scanPorts(host_ip_domain, only_open)
	},
}

func init() {
	scan_portsCmd.Flags().StringVarP(&host_ip_domain, "host", "", "", "Host's domain or ip to ping")
	scan_portsCmd.Flags().BoolVarP(&only_open,"only-open", "",false, "Show only open ports")

	if err := scan_portsCmd.MarkFlagRequired("host"); err != nil{
		fmt.Println(err)
	}

	
	NetCmd.AddCommand(scan_portsCmd)
}


