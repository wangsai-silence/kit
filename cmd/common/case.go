package common

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var upperCmd = &cobra.Command{
	Use:   "upper",
	Short: "convert to upper case",
	RunE:  upper,
}

var lowerCmd = &cobra.Command{
	Use:   "lower",
	Short: "convert to lower case",
	RunE:  lower,
}

func upper(cmd *cobra.Command, args []string) (err error) {
	if len(args) == 0 {
		err = errors.New("need one argument")
		return
	}

	fmt.Println(strings.ToUpper(args[0]))
	return
}

func lower(cmd *cobra.Command, args []string) (err error) {
	if len(args) == 0 {
		err = errors.New("need one argument")
		return
	}

	fmt.Println(strings.ToLower(args[0]))

	return
}
