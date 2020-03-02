### HMS server client push. Custom realization. Auth Client Password Mode.

- auto auth token refresher 
- sending push through hms

Documentation - [api HMS push](https://developer.huawei.com/consumer/en/doc/development/HMS-References/push-sendapi)</br>
Sample code by huawei - [push-servergosdk](https://developer.huawei.com/consumer/en/doc/development/HMS-Examples/push-servergosdk) 

#### how to use
```go
package main

import (
	"context"
	"net/http"
	"time"

	"github.com/nburmi/huawei-push/push"
	"github.com/nburmi/huawei-push/token"
)

func main() {
	cli := &http.Client{
		Timeout: time.Second,
	}

	hmsParams := token.Params{
		HTTPDoer:     cli,
		ClientID:     "Client ID(App ID)",
		ClientSecret: "Client Secret(App secret)",
	}
	
	tokener, err := token.New().SetByParams(hmsParams).Build()
	if err != nil {
		//handle error
	}

	ctx := context.Background()

	// updating the token before it expires. Clent
	tokener, err = token.NewRefresher(ctx, tokener).SetSubTime(time.Second * 5).Build()
	if err != nil {
		//handle error
	}

	pusher := push.New(hmsParams.ClientID, tokener, cli)

	//send push
	resp, err = pusher.Push(&push.Message{
		Data:   "data",
		Tokens: []string{"DEVICE TOKEN 1"},
	})

	if err != nil {
		//handle error
	}

	//check response by documentation https://developer.huawei.com/consumer/en/doc/development/HMS-References/push-sendapi
	/*
	type Response struct {
		StatusCode int    `json:"-"` //http status code
		Code      string `json:"code"`
		Message   string `json:"msg"`
		RequestID string `json:"requestId"`
	}
	*/
}

```
