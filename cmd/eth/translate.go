package eth

import (
	"fmt"
	"github.com/spf13/cobra"
	"math/big"
)

var toWeiCmd = &cobra.Command{
	Use:   "towei",
	Short: "ether to wei. example: kit eth towei 1 gwei",
	RunE:  toWei,
}

var toEtherCmd = &cobra.Command{
	Use:   "toether",
	Short: "wei to ether. example kit eth toether 10 gwei",
	RunE:  toEther,
}

var toCmd = &cobra.Command{
	Use:   "to",
	Short: "translate. example kit eth to gwei 10 wei",
	RunE:  to,
}

var unitMap = map[string]uint{
	"ether": 18,
	"gwei":  9,
	"mwei":  6,
	"kwei":  3,
	"wei":   0,
}

func to(cmd *cobra.Command, args []string) (err error) {
	target := args[0]

	unit := "ether"
	if len(args) > 2 {
		unit = args[2]
	}

	expTowei, exist := unitMap[unit]
	if !exist {
		err = fmt.Errorf("unknown unit:%v", unit)
		return
	}

	expToTarget, exist := unitMap[target]
	if !exist {
		err = fmt.Errorf("unknown unit:%v", target)
		return
	}

	if expTowei >= expToTarget {
		num, ok := big.NewInt(0).SetString(args[1], 10)
		if !ok {
			return
		}
		res := big.NewInt(0).Mul(num, big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(expTowei-expToTarget)), nil))
		fmt.Println(res.String())
	} else {
		num, ok := big.NewFloat(0).SetString(args[1])
		if !ok {
			return
		}
		res := big.NewFloat(0).Quo(num, big.NewFloat(0).SetInt(big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(expToTarget-expTowei)), nil)))
		fmt.Println(res.Text('f', -1))
	}

	return
}

func toWei(cmd *cobra.Command, args []string) (err error) {
	num, ok := big.NewInt(0).SetString(args[0], 10)
	if !ok {
		return
	}

	unit := "ether"
	if len(args) > 1 {
		unit = args[1]
	}

	exp, exist := unitMap[unit]
	if !exist {
		err = fmt.Errorf("unknown unit:%v", unit)
		return
	}

	fmt.Println(big.NewInt(0).Mul(num, big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(exp)), nil)).String())

	return
}

func toEther(cmd *cobra.Command, args []string) (err error) {
	num, ok := big.NewFloat(0).SetString(args[0])
	if !ok {
		return
	}

	unit := "wei"
	if len(args) > 1 {
		unit = args[1]
	}

	exp, exist := unitMap[unit]
	if !exist {
		err = fmt.Errorf("unknown unit:%v", unit)
		return
	}

	res := big.NewFloat(0).Quo(num, big.NewFloat(0).SetInt(big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(18-exp)), nil)))
	fmt.Println(res.Text('f', -1))

	return
}
