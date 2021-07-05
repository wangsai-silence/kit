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
	rpcCmd.AddCommand(blockHeightCmd)
	rpcCmd.AddCommand(txCmd)

	EthCmd.AddCommand(decodeCmd)
	EthCmd.AddCommand(rpcCmd)

	EthCmd.AddCommand(toWeiCmd)
	EthCmd.AddCommand(toEtherCmd)
}
