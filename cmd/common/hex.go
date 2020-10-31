package common

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/spf13/cobra"
)

var unhexCmd = &cobra.Command{
	Use:   "unhex",
	Short: "unhex int from 16 to 10",
	RunE:  unhex,
}
var hexCmd = &cobra.Command{
	Use:   "hex",
	Short: "hex int from 10 to 16",
	RunE:  hex,
}

func unhex(cmd *cobra.Command, args []string) (err error) {
	if len(args) == 0 {
		err = errors.New("need one argument")
		return
	}

	big, err := hexutil.DecodeBig(args[0])
	if err != nil {
		return
	}

	fmt.Println(big.String())

	return
}

func hex(cmd *cobra.Command, args []string) (err error) {
	if len(args) == 0 {
		err = errors.New("need one argument")
		return
	}

	big, passed := big.NewInt(0).SetString(args[0], 10)
	if !passed {
		err = fmt.Errorf("input %v is not a valid integer", args[0])
		return
	}

	fmt.Println(hexutil.EncodeBig(big))

	return
}
