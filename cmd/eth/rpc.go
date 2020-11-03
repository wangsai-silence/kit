package eth

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cobra"
)

var rpcCmd = &cobra.Command{
	Use:   "rpc",
	Short: "get info via rpc",
	RunE:  rpc,
}

const infuraHost = "https://mainnet.infura.io/v3/c089f8ed1d7344f38393cee5426287b8"

type Response struct {
	Id      int             `json:"id"`
	Jsonrpc string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result"`
}

func rpc(cmd *cobra.Command, args []string) (err error) {
	if len(args) == 0 {
		err = errors.New("need at least one argument")
		return
	}

	method, paramStr := args[0], args[1:]

	params := make([]interface{}, 0, len(paramStr))
	for _, param := range paramStr {
		params = append(params, param)
	}

	resp, err := innerRpc(cmd, method, params)

	if err != nil {
		return
	}

	fmt.Println(string(resp.Result))

	return
}

func innerRpc(cmd *cobra.Command, method string, params []interface{}) (response *Response, err error) {
	if params == nil {
		params = []interface{}{}
	}

	body, err := json.Marshal(struct {
		Jsonrpc string        `json:"jsonrpc"`
		Method  string        `json:"method"`
		Params  []interface{} `json:"params"`
		Id      int           `json:"id"`
	}{
		Jsonrpc: "2.0",
		Id:      1,
		Method:  method,
		Params:  params,
	})

	if err != nil {
		return
	}

	host, err := cmd.Flags().GetString("host")
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", host, bytes.NewBuffer(body))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	response = &Response{}
	err = json.Unmarshal(bodyBytes, response)
	if err != nil {
		return
	}

	return
}

var blockHeightCmd = &cobra.Command{
	Use:   "height",
	Short: "get current block number via rpc",
	RunE:  height,
}

func height(cmd *cobra.Command, args []string) (err error) {
	resp, err := innerRpc(cmd, "eth_blockNumber", nil)
	if err != nil {
		return
	}

	big, err := hexutil.DecodeBig(string(resp.Result[1 : len(resp.Result)-1]))
	if err != nil {
		return
	}

	fmt.Println(big.String())
	return
}

var txCmd = &cobra.Command{
	Use:   "tx",
	Short: "get tx info",
	RunE:  tx,
}

func tx(cmd *cobra.Command, args []string) (err error) {
	if len(args) == 0 {
		err = errors.New("need tx hash as argument")
		return
	}

	resp, err := innerRpc(cmd, "eth_getTransactionByHash", []interface{}{args[0]})
	if err != nil {
		return
	}

	org := &types.Transaction{}
	err = json.Unmarshal(resp.Result, org)
	if err != nil {
		return
	}

	wrapper, err := WrapTransaction(org)
	if err != nil {
		return
	}

	data, err := json.MarshalIndent(wrapper, "", "    ")
	if err != nil {
		return
	}

	fmt.Println(string(data))
	return
}
