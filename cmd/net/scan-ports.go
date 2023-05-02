package net

import (
	// "fmt"
	"fmt"

	"github.com/spf13/cobra"
)

var only_open bool
var only_tcp bool
var only_udp bool
var start_port int
var end_port int

var scan_portsCmd = &cobra.Command{
	Use:   "scan-ports",
	Short: "Network utility that probes and identifies open ports on a target system.",
	Long: `Port scanning work by sending various types of packets to a range of ports on a target system and monitoring the responses to determine which ports are open, closed, or filtered.`,
	Run: func(cmd *cobra.Command, args []string) {
		if (start_port > end_port){
			fmt.Printf("Starting port cannot be bigger than end port")
			return
		}
		do_scanPorts(host_ip_domain, only_open, only_tcp, only_udp, start_port, end_port)
	},
}

func init() {
	scan_portsCmd.Flags().StringVarP(&host_ip_domain, "host", "", "", "Host's domain or ip to ping")
	scan_portsCmd.Flags().BoolVarP(&only_open,"only-open", "",false, "Show only open ports")

	scan_portsCmd.Flags().BoolVarP(&only_tcp, "only-tcp", "", false, "Show only TCP scan")
	scan_portsCmd.Flags().BoolVarP(&only_udp, "only-udp", "", false, "Show only UDP scan")
	scan_portsCmd.MarkFlagsMutuallyExclusive("only-tcp", "only-udp")

	scan_portsCmd.Flags().IntVarP(&start_port, "start-port", "", 0, "Start port")
	scan_portsCmd.Flags().IntVarP(&end_port, "end-port", "", 0, "End port")


	scan_portsCmd.MarkFlagRequired("host")
	scan_portsCmd.MarkFlagRequired("start-port")
	scan_portsCmd.MarkFlagRequired("end-port")

	
	NetCmd.AddCommand(scan_portsCmd)
}


