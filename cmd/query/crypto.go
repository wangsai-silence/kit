package query

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

var crypto = &cobra.Command{
	Use:   "crypto",
	Short: "crypto price from exchanges. ex: crypto btcusdt(default) huobi(default)",
	RunE:  price,
}

func price(cmd *cobra.Command, args []string) (err error) {
	symbol := "btcusdt"
	exchange := "huobi"

	if len(args) > 0 {
		symbol = args[0]
	}
	if len(args) > 1 {
		exchange = args[1]
	}
	url := fmt.Sprintf("https://addons.wangsai.site/getPrice?exchange=%s&symbol=%s", exchange, symbol)
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		return
	}

	bData, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	res := make(map[string]interface{}, 0)
	err = json.Unmarshal(bData, &res)
	if err != nil {
		return
	}

	fmt.Println(res["data"])

	return
}
