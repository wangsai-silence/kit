package common

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var lenCmd = &cobra.Command{
	Use:   "len",
	Short: "calc length for string",
	RunE:  length,
}

func length(cmd *cobra.Command, args []string) (err error) {
	if len(args) == 0 {
		err = errors.New("need one argument")
		return
	}

	fmt.Println(len(args[0]))

	return
}
