package push

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/nburmi/huawei-push/token"
)

type mockHTTPDoer struct {
	url         string
	contentType string
	body        []byte
}

type bf struct {
	bytes.Buffer
}

func (b *bf) Close() error {
	return nil
}

func (m *mockHTTPDoer) Do(req *http.Request) (resp *http.Response, err error) {
	m.url = req.URL.String()
	m.contentType = req.Header.Get("Content-Type")
	m.body, _ = ioutil.ReadAll(req.Body)

	var bf bf
	bf.WriteString(`{
		"access_token": "ACCESS_TOKEN",
		"expires_in": 3600,
		"token_type": "Bearer"
	}`)

	resp = &http.Response{
		Body: &bf,
	}

	return resp, nil
}
func TestPusher(t *testing.T) {
	cli := &mockHTTPDoer{}

	tokener, err := token.New().
		SetByParams(token.Params{
			HTTPDoer:     cli,
			ClientID:     "APP ID",
			ClientSecret: "APP SECRET"}).
		Build()
	if err != nil {
		t.Error(err)
	}

	tokener, err = token.NewRefresher(context.Background(), tokener).
		SetSubTime(time.Second * 5).
		Build()
	if err != nil {
		t.Error(err)
	}

	p := New("APP ID", tokener, cli)
	_, err = p.PushValidate(&Message{
		Data:   "{\"msg\":\"this is message\", \"title\":\"simple title\"}",
		Tokens: []string{"DEVICE TOKEN"},
	})

	if err != nil {
		t.Error(err)
	}

	expected := `{"message":{"data":"{\"msg\":\"this is message\", \"title\":\"simple title\"}","token":["DEVICE TOKEN"]},"validate_only":true}`
	if expected != string(cli.body) {
		t.Errorf("body not equal, expected %s, got %s", expected, cli.body)
	}
}
