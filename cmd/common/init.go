package common

import "github.com/spf13/cobra"

var CommonCmds = []*cobra.Command{
	hexCmd,
	unhexCmd,
	unixCmd,
}

func init() {
}
