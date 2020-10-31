package eth

import (
	"github.com/spf13/cobra"
)

var EthCmd = &cobra.Command{
	Use:   "eth",
	Short: "eth tools",
}

func init() {
	rpcCmd.PersistentFlags().StringP("host", "o", infuraHost, "set host name")

	EthCmd.AddCommand(decodeCmd)
	EthCmd.AddCommand(rpcCmd)
}
