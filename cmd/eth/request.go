package eth

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

var rpcCmd = &cobra.Command{
	Use:   "rpc",
	Short: "get info via rpc",
	RunE:  rpc,
}

const infuraHost = "https://mainnet.infura.io/v3/c089f8ed1d7344f38393cee5426287b8"

func rpc(cmd *cobra.Command, args []string) (err error) {
	if len(args) == 0 {
		err = errors.New("need at least one argument")
		return
	}

	method, params := args[0], args[1:]

	body, err := json.Marshal(struct {
		Jsonrpc string   `json:"jsonrpc"`
		Method  string   `json:"method"`
		Params  []string `json:"params"`
		Id      int      `json:"id"`
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

	fmt.Println(string(bodyBytes))

	return
}
