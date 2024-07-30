package client

import (
	"context"
	"crypto/tls"
	"github.com/GoPlusSecurity/goplus-usm/util/ecode"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
	"time"
)

const TARGET_ADDR = "usm.gopluslabs.io"

type USMClient struct {
	conn         *grpc.ClientConn
	clientOption clientOption
}

// create grpc client
func (u *USMClient) newDetectConn(opts ...ClientOption) error {
	for _, opt := range opts {
		opt.apply(&u.clientOption)
	}

	// keepalive
	var kacp = keepalive.ClientParameters{
		Time:                20 * time.Second, // send pings every 20s seconds if there is no activity
		Timeout:             60 * time.Second, // wait 60 second for ping ack before considering the connection dead
		PermitWithoutStream: true,             // send pings even without active streams
	}

	// retry
	var retryConfig = `{
		"methodConfig": [{
		  "name": [{"service": "detect.service.v1.Detect","method":"DetectTx"}],
		  "retryPolicy": {
			  "MaxAttempts": 3,
			  "InitialBackoff": "0.5s",
			  "MaxBackoff": "1s",
			  "BackoffMultiplier": 2,
			  "RetryableStatusCodes": [ "UNAVAILABLE","RESOURCE_EXHAUSTED" ]
		  }
		}]}`

	conn, err := grpc.NewClient(TARGET_ADDR,
		grpc.WithDefaultServiceConfig(retryConfig),
		grpc.WithKeepaliveParams(kacp),
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: true,
		})))
	if err != nil {
		return ecode.GenInternalError(ecode.DetectRpcServerConnErr, err)
	}
	u.conn = conn
	return nil
}

func (u *USMClient) genGrpcCTX(ctx context.Context) context.Context {
	if u.clientOption.timeOut > 0 {
		ctx, _ = context.WithTimeout(ctx, u.clientOption.timeOut)
	}
	if u.clientOption.key != "" {
		ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("x-api-key", u.clientOption.key))
	}
	return ctx
}

func (u *USMClient) Close() error {
	if u.conn != nil {
		return u.conn.Close()
	}
	return nil
}
