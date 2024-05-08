package query

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

var crypto = &cobra.Command{
	Use:     "crypto",
	Aliases: []string{"c"},
	Short:   "crypto price from exchanges. ex: crypto btcusdt(default) huobi(default)",
	RunE:    price,
}

type style struct {
	format    string
	uppercase bool
}

var symbolStyles = map[string]*style{
	"huobi": &style{
		format:    "%s%s",
		uppercase: false,
	},
	"binance": &style{
		format:    "%s%s",
		uppercase: true,
	},
	"coinbase": &style{
		format:    "%s-%s",
		uppercase: true,
	},
	"kucoin": &style{
		format:    "%s-%s",
		uppercase: true,
	},
}

var baseTokens = []string{"usdt", "usdc", "busd", "usd", "btc", "eth", "bnb", "ht"}

func formatSymbol(symbol string, exchange string) (formatSymbol string, err error) {
	formatSymbol = symbol
	symbolStyle := symbolStyles[strings.ToLower(exchange)]
	if symbolStyle == nil {
		return
	}

	symbol = strings.ToLower(symbol)
	var baseToken string
	for _, base := range baseTokens {
		if strings.HasSuffix(symbol, base) {
			baseToken = base
			break
		}
	}

	if baseToken == "" {
		return
	}

	formatSymbol = fmt.Sprintf(symbolStyle.format, strings.TrimSuffix(symbol, baseToken), baseToken)
	if symbolStyle.uppercase {
		formatSymbol = strings.ToUpper(formatSymbol)
	}

	return
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

	symbol, err = formatSymbol(symbol, exchange)
	if err != nil {
		return
	}
	url := fmt.Sprintf("https://addons.wangsai.cloud/getPrice?exchange=%s&symbol=%s", exchange, symbol)
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
