package eth

import "github.com/ethereum/go-ethereum/core/types"

type TxWrapper struct {
	Origin   *types.Transaction `json:"origin"`
	From     string             `json:"from"`
	Amount   string             `json:"amount"`
	ChainId  int64              `json:"chainId"`
	Gas      uint64             `json:"gas"`
	GasPrice string             `json:"gasePrice"`
	Nonce    uint64             `json:"nonce"`
}

func WrapTransaction(tx *types.Transaction) (wrapper *TxWrapper, err error) {
	var signer types.Signer
	if tx.Protected() {
		signer = types.NewEIP155Signer(tx.ChainId())
	} else {
		signer = &types.HomesteadSigner{}
	}

	sender, err := types.Sender(signer, tx)
	if err != nil {
		return
	}

	wrapper = &TxWrapper{
		Origin:   tx,
		ChainId:  tx.ChainId().Int64(),
		From:     sender.String(),
		Amount:   tx.Value().String(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice().String(),
		Nonce:    tx.Nonce(),
	}

	return
}
