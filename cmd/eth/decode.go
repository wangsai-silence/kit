package eth

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cobra"
)

var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "decode raw transaction",
	RunE:  decode,
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
	err = tx.UnmarshalBinary(bytes)
	if err != nil {
		return
	}

	wrapper, err := WrapTransaction(tx)
	if err != nil {
		return
	}

	data, err := json.MarshalIndent(wrapper, "", "    ")
	if err != nil {
		return
	}

	fmt.Println(string(data))

	return
}
