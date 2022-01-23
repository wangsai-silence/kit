package query

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/djimenez/iconv-go"
	"github.com/spf13/cobra"
)

var stock = &cobra.Command{
	Use:   "stock",
	Short: "stock price for shanghai/shenzhen a share market. ex: stock sh000001(default)",
	RunE:  stockPrice,
}

type StockResponse struct {
	Code string        `json:"code"`
	Date int64         `json:"date"`
	Snap []interface{} `json:"snap"`
}

func stockPrice(cmd *cobra.Command, args []string) (err error) {
	code := "sh000001"

	if len(args) > 0 {
		code = args[0]
	}

	url := fmt.Sprintf("https://qt.gtimg.cn/q=%s", code)
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		return
	}

	bData, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	datas := strings.Split(string(bData), "~")
	if len(datas) < 3 {
		return fmt.Errorf("unexpect response:%v", string(bData))
	}
	//
	//datas := strings.Split(sp[1], ",")
	//if len(datas) < 4 {
	//	return fmt.Errorf("unexpect response:%v", string(bData))
	//}

	name, err := iconv.ConvertString(datas[1], "GBK", "utf-8")
	if err != nil {
		return
	}
	prev, err := strconv.ParseFloat(datas[4], 64)
	if err != nil {
		return
	}
	latest, err := strconv.ParseFloat(datas[3], 64)
	if err != nil {
		return
	}

	fmt.Println(fmt.Sprintf(`
	name	  : %v
	prev close: %v
	latest    : %v
	rate      : %.2f %%
	`, name, prev, latest, (latest-prev)*100/prev))

	return
}
