package main

import (
	"github.com/kit/cmd/common"
	"github.com/kit/cmd/eth"
	"github.com/kit/cmd/query"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tool",
	Short: "A tool set for personal use",
	Long:  ``,
}

func main() {
	rootCmd.AddCommand(common.CommonCmds...)
	rootCmd.AddCommand(eth.EthCmd, query.QueryCmd)

	rootCmd.Execute()
}
