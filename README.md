# GoPlus User Security Model
For building your own applications and integrating your platforms with GoPlus User Security Model seamlessly, you can use GoPlus User Security Model SDKs. We currently provide SDKs in the following languages:

## How to use goplus-usm

### create evm client
```go
USMEVMclient, err := client.NewUSMEVMClient(
    WithKey("This is a key"),
    WithTimeOut(10*time.Second),
)
defer USMEVMclient.Close()
```

### send transaction by DetectRequest
```go
var req = client.DetectRequest{
    From:     "0xbE2d195D57217941fAb5bC8B554ad60899e99a0F",
    To:       "0x97E542Ec6B81Dea28F212775Ce8Ac436Ab77a7df",
    Gas:      21000,
    GasPrice: "75451731683",
    Value:    "1000000000000000000",
    Data:     "0x629c3fc3",
    ChainID:  "1",
    Nonce:    1, 
    Hash:   "0xe60bd022c8a330cb9e2aed0c1fbb7a15097d78c9d6cca0254b42d653dbf01575"
}
err, resp := USMEVMclient.DialDetect(req)
if err != nil {
    panic(err)
}
```

### send transaction by raw transaction
```go
var rawTX = "0xf86e01851191460ee38252089497e542ec6b81dea28f212775ce8ac436ab77a7df880de0b6b3a764000082307826a0963edb6e57c2c6bd0d4a8d827a53f6f9e164f09dd1bc5d1f8580c020abad56b5a04343639754d9d9662e1bb6cdeee65f5c315d27fe183119763c9c235876a3d2f8"
err, resp := USMEVMclient.DialDetectWithRawTX(rawTX)
if err != nil {
    panic(err)
}
```