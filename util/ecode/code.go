package ecode

var (
	OK = New("ok", "ok")

	ServerErr                  = New("SERVER_ERR", "Server error")
	ReqFormatErr               = New("REQ_FORMAT_ERR", "request format error")
	DetectRpcServerConnErr     = New("DETECT_RPC_SERVER_CONN_ERR", "detect rpc server conn error,please call administrator as soon as possible")
	DetectRpcServerInternalErr = New("DETECT_RPC_SERVER_INTERNAL_ERROR", "detect rpc server internal error")
	DecodeRawTransactionErr    = New("DECODE_RAW_TRANSACTION_ERROR", "decode raw transaction error")
)
