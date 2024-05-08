package common

import "github.com/spf13/cobra"

var CommonCmds = []*cobra.Command{
	hexCmd,
	b64Cmd,
	unixCmd,
	strCmd,
}
