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
	signers := []types.Signer{
		types.NewLondonSigner(tx.ChainId()),
		types.NewEIP2930Signer(tx.ChainId()),
		types.NewEIP155Signer(tx.ChainId()),
		types.HomesteadSigner{},
		types.FrontierSigner{},
	}

	for _, signer := range signers {
		sender, err1 := types.Sender(signer, tx)
		if err1 != nil {
			err = err1
			continue
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
	}

	if wrapper != nil {
		err = nil
	}

	return
}
