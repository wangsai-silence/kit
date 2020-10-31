package eth

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/spf13/cobra"
)

var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "decode raw transaction",
	RunE:  decode,
}

type Wrapper struct {
	Origin  *types.Transaction `json:"origin"`
	From    string             `json:"from"`
	ChainId int64              `json:"chainId"`
}

func decode(cmd *cobra.Command, args []string) (err error) {
	if len(args) == 0 {
		err = errors.New("need one argument")
		return
	}

	bytes, err := hexutil.Decode(args[0])
	if err != nil {
		return
	}

	tx := new(types.Transaction)
	err = rlp.DecodeBytes(bytes, tx)
	if err != nil {
		return
	}

	sender, err := types.Sender(types.NewEIP155Signer(tx.ChainId()), tx)
	if err != nil {
		return
	}

	wrapper := &Wrapper{
		Origin:  tx,
		From:    sender.String(),
		ChainId: tx.ChainId().Int64(),
	}

	data, err := json.MarshalIndent(wrapper, "", "    ")
	if err != nil {
		return
	}

	fmt.Println(string(data))

	return
}
