package cmd

import (
	"os"
	"github.com/spf13/cobra"
	"github.com/theveloper-pl/hydra/cmd/net"
)


var rootCmd = &cobra.Command{
	Use:   "hydra",
	Short: "A Go-based networking project.",
	Long: `The Network Toolbox project in Go is an advanced 
suite of tools designed for IT professionals, 
network administrators, and security enthusiasts to analyze and monitor networks. 
This project includes a variety of features that help diagnose network issues, 
analyze security, and monitor internet connections.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubcommandPalettes(){
	rootCmd.AddCommand(net.NetCmd)
}


func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	addSubcommandPalettes()
}


