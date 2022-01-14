package query

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

var stock = &cobra.Command{
	Use:   "stock",
	Short: "stock price for shanghai/shenzhen a share market. ex: stock 000001(default)",
	RunE:  stockPrice,
}

type StockResponse struct {
	Code string        `json:"code"`
	Date int64         `json:"date"`
	Snap []interface{} `json:"snap"`
}

func stockPrice(cmd *cobra.Command, args []string) (err error) {
	code := "000001"

	if len(args) > 0 {
		code = args[0]
	}

	url := fmt.Sprintf("http://yunhq.sse.com.cn:32041/v1/sh1/snap/%v?select=name,prev_close,last,chg_rate", code)
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		return
	}

	bData, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	res := &StockResponse{}
	err = json.Unmarshal(bData, &res)
	if err != nil {
		return
	}

	fmt.Println(fmt.Sprintf(`
	prev close: %v
	latest    : %v
	rate      : %v %%
	`, res.Snap[1], res.Snap[2], res.Snap[3]))

	return
}
