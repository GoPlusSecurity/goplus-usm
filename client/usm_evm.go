package client

import (
	"context"
	"encoding/hex"
	pb "github.com/GoPlusSecurity/goplus-usm/api"
	"github.com/GoPlusSecurity/goplus-usm/util"
	"github.com/GoPlusSecurity/goplus-usm/util/ecode"
	"strings"
	"time"
)

const USM_EVM_CLIENT_DEFAULT_TIMEOUT = 10 * time.Second

type USMEVMClient struct {
	USMClient
}

func NewUSMEVMClient(opts ...ClientOption) (*USMEVMClient, error) {
	var client = USMEVMClient{}
	client.setDefaultTimeOut()
	err := client.newDetectConn(opts...)
	return &client, err
}

func (u *USMEVMClient) setDefaultTimeOut() {
	u.clientOption.timeOut = USM_EVM_CLIENT_DEFAULT_TIMEOUT
}

// DialDetect detect transaction by DetectEVMRequest
func (u *USMEVMClient) DialDetect(ctx context.Context, detectReq *pb.DetectTxRequest) (res *pb.DetectTxResponse, err error) {
	if !strings.HasPrefix(detectReq.Data, "0x") {
		detectReq.Data = "0x" + detectReq.Data
	}
	res, err = pb.NewDetectClient(u.conn).DetectTx(u.genGrpcCTX(ctx), detectReq)
	if err != nil {
		return res, ecode.GenInternalError(ecode.DetectRpcServerInternalErr, err)
	}
	return res, nil
}

// DialDetectWithRawTX detect transaction by raw transaction
func (u *USMEVMClient) DialDetectWithRawTX(ctx context.Context, rawTX string) (res *pb.DetectTxResponse, err error) {
	from, tx, err := util.DecodeRawTX(rawTX)
	if err != nil {
		return res, ecode.GenInternalError(ecode.DecodeRawTransactionErr, err)
	}

	data := hex.EncodeToString(tx.Data())
	var req = pb.DetectTxRequest{
		From:     from,
		To:       tx.To().Hex(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice().String(),
		Value:    tx.Value().String(),
		Data:     data,
		ChainId:  tx.ChainId().String(),
		Nonce:    tx.Nonce(),
		Hash:     tx.Hash().Hex(),
	}
	return u.DialDetect(ctx, &req)
}
