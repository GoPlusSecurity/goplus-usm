package util

import (
	"bytes"
	"github.com/ethereum/go-ethereum/core/types"
	"testing"
)

func TestMul(t *testing.T) {
	tx := types.Transaction{}
	cases := []struct {
		Name         string
		RawTX        string
		ExpectedTX   *types.Transaction
		ExpectedFrom string
	}{
		{"null rawTX", "", &tx, ""},
		{"wrong rawTX_01", "0x2dadias8dasd", &tx, ""},
		{"wrong rawTX_02", "851191460ee38252089497e542ec6b81dea28f212775ce8ac4", &tx, ""},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			from, transaction, err := DecodeRawTX(c.RawTX)
			if err != nil {
				t.Log("err:", err.Error())
			}
			t.Log("from:", from, "transaction:", transaction)
			if *transaction != *c.ExpectedTX || from != c.ExpectedFrom {
				t.Error("test fail")
			}
		})
	}

	rawTX := "0xf86e01851191460ee38252089497e542ec6b81dea28f212775ce8ac436ab77a7df880de0b6b3a764000082307826a0963edb6e57c2c6bd0d4a8d827a53f6f9e164f09dd1bc5d1f8580c020abad56b5a04343639754d9d9662e1bb6cdeee65f5c315d27fe183119763c9c235876a3d2f8"
	from, transaction, err := DecodeRawTX(rawTX)
	if err != nil {
		t.Log("err:", err.Error())
	}
	t.Log("from:", from, "transaction:", transaction)
	if transaction == nil || from == "" {
		t.Error("test fail")
	}
	if transaction.Nonce() != 1 || bytes.Compare(transaction.Data(), []byte("0x")) != 0 ||
		transaction.To().Hex() != "0x97E542Ec6B81Dea28F212775Ce8Ac436Ab77a7df" || transaction.Gas() != 21000 ||
		transaction.GasPrice().Int64() != 75451731683 || transaction.Value().Int64() != 1000000000000000000 ||
		transaction.ChainId().String() != "1" || from != "0xbE2d195D57217941fAb5bC8B554ad60899e99a0F" {
		t.Error("test fail")
	}
}
