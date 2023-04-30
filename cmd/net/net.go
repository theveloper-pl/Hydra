package net

import (

	"github.com/spf13/cobra"
)

var (
	host_ip_domain string
	pings_number int
)

var NetCmd = &cobra.Command{
	Use:   "net",
	Short: "net submodule of Hydra",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

}
