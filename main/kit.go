package main

import (
	"github.com/kit/cmd/common"
	"github.com/kit/cmd/eth"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tool",
	Short: "A tool set for personal use",
	Long:  ``,
}

func main() {
	for _, cmd := range common.CommonCmds {
		rootCmd.AddCommand((cmd))
	}
	rootCmd.AddCommand(eth.EthCmd)

	rootCmd.Execute()
}
