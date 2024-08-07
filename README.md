# GoPlus User Security Model
For building your own applications and integrating your platforms with GoPlus User Security Model seamlessly, you can use GoPlus User Security Model SDKs. We currently provide SDKs in the following languages:

## How to use goplus-usm

### Installation
To start, add goplus-usm as a dependency to your project:
```go
go get github.com/GoPlusSecurity/goplus-usm
go mod tidy
```

### create evm client
```go
import (
    "github.com/GoPlusSecurity/goplus-usm/client"
    "time"
)

USMEVMclient, err := client.NewUSMEVMClient(
    client.WithKey("This is a key"),
    client.WithTimeOut(10*time.Second),
)
defer USMEVMclient.Close()
if err != nil {
    panic(err)
}
```

### send transaction by DetectRequest
```go
import (
    "context"
    "fmt"
    gapi "github.com/GoPlusSecurity/goplus-usm/api"
    "github.com/GoPlusSecurity/goplus-usm/client"
    "time"
)

// IsIntercept: false
var req = gapi.DetectTxRequest{
    From:     "0xbE2d195D57217941fAb5bC8B554ad60899e99a0F",
    To:       "0x97E542Ec6B81Dea28F212775Ce8Ac436Ab77a7df",
    Gas:      21000,
    GasPrice: "75451731683",
    Value:    "1000000000000000000",
    Data:     "0x629c3fc3",
    ChainId:  "1",
    Nonce:    1,
    Hash:     "0xe60bd022c8a330cb9e2aed0c1fbb7a15097d78c9d6cca0254b42d653dbf01575",
}
resp, err := USMEVMclient.DialDetect(context.Background(), &req)
if err != nil {
    panic(err)
}
fmt.Println("IsIntercept: ", resp.IsIntercept)
fmt.Println("RiskInfo: ", resp.RiskInfo)
```

### send transaction by raw transaction
```go
import (
    "context"
    "fmt"
    gapi "github.com/GoPlusSecurity/goplus-usm/api"
    "github.com/GoPlusSecurity/goplus-usm/client"
    "time"
)

// IsIntercept: true
var rawTX = "0xf86c02851191460ee382752f94f6263db8371a633ee97793dfe26c533a1edd3920880de0b6b3a76400008025a0a7e477feb1c695c6d84abe12997078dae4679cdde21e6411c57427ef22d22236a02ec2f186ffe26523c3ec170b24fd3f8bf5b813e41f1f2e2a10c866e3c7df7f84"
resp, err = USMEVMclient.DialDetectWithRawTX(context.Background(), rawTX)
if err != nil {
    panic(err)
}
fmt.Println("IsIntercept: ", resp.IsIntercept)
fmt.Println("RiskInfo: ", resp.RiskInfo)
```