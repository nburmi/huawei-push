### HMS server client push. Custom realization.

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

	hms "github.com/nburmi/huawei-push"
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

	// updating the token before it expires.
	tokener, err = token.NewRefresher(ctx, tokener).SetSubTime(time.Second * 5).Build()
	if err != nil {
		//handle error
	}

	// check response and create error
	tokener = token.NewCheckTokener(tokener)
	if err != nil {
		//handle error
	}

	pusher := push.New(hmsParams.ClientID, tokener, cli)

	//for check response and return error
	pusher = push.NewResponseChecker(pusher)

	//send push
	_, err = pusher.Push(&push.Message{
		Data:   "data",
		Tokens: []string{"DEVICE TOKEN 1"},
	})

	if err != nil {
		//handle error
	}
}

```