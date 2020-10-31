package main

import (
	"github.com/spf13/cobra"
	"github.com/tools/cmd/common"
	"github.com/tools/cmd/eth"
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
