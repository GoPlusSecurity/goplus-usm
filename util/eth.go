package util

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strings"
)

func DecodeRawTX(rawTX string) (from string, tx *types.Transaction, err error) {
	rawTX = strings.TrimPrefix(rawTX, "0x")

	tx = &types.Transaction{}
	rawTxBytes, err := hex.DecodeString(rawTX)
	if err != nil {
		return
	}

	err = tx.UnmarshalBinary(rawTxBytes)
	if err != nil {
		return
	}

	signer := types.NewCancunSigner(big.NewInt(tx.ChainId().Int64()))
	address, err := types.Sender(signer, tx)
	if err == nil {
		from = address.String()
	} else {
		return
	}
	return
}
