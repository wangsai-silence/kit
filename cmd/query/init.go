package query

import (
	"github.com/spf13/cobra"
)

var QueryCmd = &cobra.Command{
	Use:   "query",
	Short: "query info via http",
}

func init() {
	QueryCmd.AddCommand(crypto)
	QueryCmd.AddCommand(stock)
}
