package common

import (
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	b64Cmd.AddCommand(b64DecodeCmd)
}

var b64Cmd = &cobra.Command{
	Use:     "base64",
	Aliases: []string{"b64"},
	Short:   "convert string between utf8 and base64",
}
var b64DecodeCmd = &cobra.Command{
	Use:     "decode",
	Aliases: []string{"d", "de"},
	Short:   "decode base64 data to string",
	RunE:    decode,
}

func decode(cmd *cobra.Command, args []string) (err error) {
	if len(args) == 0 {
		err = errors.New("need one argument")
		return
	}

	data, err := base64.StdEncoding.DecodeString(args[0])
	if err != nil {
		data, err = base64.URLEncoding.DecodeString(args[0])
		if err != nil {
			return
		}
		fmt.Println("decode with base64 url:")
	}

	fmt.Println(string(data))

	return
}
