package common

import "github.com/spf13/cobra"

var CommonCmds = []*cobra.Command{
	hexCmd,
	unhexCmd,
	unixCmd,
	upperCmd,
	lowerCmd,
	lenCmd,
	b64Cmd,
}

func init() {
}
