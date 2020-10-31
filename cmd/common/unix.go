package common

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var unixCmd = &cobra.Command{
	Use:   "unix",
	Short: "get unix timestamp",
	RunE:  getUnixTime,
}

func getUnixTime(cmd *cobra.Command, args []string) (err error) {
	fmt.Println(time.Now().Unix())

	return
}
