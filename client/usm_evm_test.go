package client

import (
	"context"
	pb "github.com/GoPlusSecurity/goplus-usm/api"
	"testing"
	"time"
)

func TestUSMEVMClient_DialDetect(t *testing.T) {
	cases := []struct {
		Name     string
		Req      pb.DetectTxRequest
		Expected pb.DetectTxResponse
	}{
		{"null req", pb.DetectTxRequest{}, pb.DetectTxResponse{IsIntercept: false}},
		{"param wrong req", pb.DetectTxRequest{
			From:     "0xbE2d195D57217941fAb5bC8B554ad60899e99a0F",
			To:       "0x97E542Ec6B81Dea28F212775Ce8Ac436Ab77a7df",
			Gas:      21000,
			GasPrice: "8740459919",
			Value:    "0",
			Data:     "0x",
			ChainId:  "1",
			Nonce:    1,
			Hash:     "",
		}, pb.DetectTxResponse{IsIntercept: false}},
		{"correct req", pb.DetectTxRequest{
			From:     "0xbE2d195D57217941fAb5bC8B554ad60899e99a0F",
			To:       "0x97E542Ec6B81Dea28F212775Ce8Ac436Ab77a7df",
			Gas:      77386,
			GasPrice: "8740459919",
			Value:    "0",
			Data:     "0x629c3fc3",
			ChainId:  "1",
			Nonce:    1,
			Hash:     "",
		}, pb.DetectTxResponse{IsIntercept: false}},
		{"intercept req", pb.DetectTxRequest{
			From:     "0xbE2d195D57217941fAb5bC8B554ad60899e99a0F",
			To:       "0xf6263db8371a633ee97793dfe26c533a1edd3920",
			Gas:      21000,
			GasPrice: "75451731683",
			Data:     "0x",
			Value:    "1000000000000000000",
			ChainId:  "1",
			Nonce:    1,
			Hash:     "",
		}, pb.DetectTxResponse{IsIntercept: true}},
	}

	client, err := NewUSMEVMClient(
		WithKey("This is a key"),
		WithTimeOut(30*time.Second),
	)
	defer client.Close()
	if err == nil {
		for _, c := range cases {
			t.Run(c.Name, func(t *testing.T) {
				resp, err := client.DialDetect(context.Background(), &c.Req)
				if err != nil {
					t.Errorf("err:%v", err)
					return
				}
				if resp.IsIntercept != c.Expected.IsIntercept {
					t.Error("test fail")
				}
			})
		}
	} else {
		t.Error(err.Error())
	}

}
