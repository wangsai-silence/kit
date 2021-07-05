package eth

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/spf13/cobra"
)

var toWeiCmd = &cobra.Command{
	Use:   "towei",
	Short: "ether to wei",
	RunE:  toWei,
}

var toEtherCmd = &cobra.Command{
	Use:   "toether",
	Short: "wei to ether",
	RunE:  toEther,
}

var unitMap = map[string]uint{
	"ether": 18,
	"gwei":  9,
	"mwei":  6,
	"kwei":  3,
	"wei":   0,
}

func toWei(cmd *cobra.Command, args []string) (err error) {
	num, err := strconv.Atoi(args[0])
	if err != nil {
		return
	}

	unit := "ether"
	if len(args) > 1 {
		unit = args[1]
	}

	exp, exist := unitMap[unit]
	if !exist {
		err = fmt.Errorf("Unknown unit:%s", unit)
		return
	}

	fmt.Println(big.NewInt(0).Mul(big.NewInt(int64(num)), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(exp)), nil)).String())

	return
}

func toEther(cmd *cobra.Command, args []string) (err error) {
	num, err := strconv.Atoi(args[0])
	if err != nil {
		return
	}

	unit := "wei"
	if len(args) > 1 {
		unit = args[1]
	}

	exp, exist := unitMap[unit]
	if !exist {
		err = fmt.Errorf("Unknown unit:%s", unit)
		return
	}

	fmt.Println(big.NewFloat(0).Quo(big.NewFloat(float64(num)), big.NewFloat(0).SetInt(big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(18-exp)), nil))).String())

	return
}
