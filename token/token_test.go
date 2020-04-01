package token

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
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

func TestGetToken(t *testing.T) {
	p := &mockHTTPDoer{}
	tokener, err := New().SetByParams(&Params{
		ClientID:     "AppID",
		ClientSecret: "AppSecret",
		HTTPDoer:     p,
	}).Build()

	if err != nil {
		t.Error(err)
	}

	token, err := tokener.Get()
	if err != nil {
		t.Error(err)
	}

	expectedBody := "client_id=AppID&client_secret=AppSecret&grant_type=client_credentials"

	switch {
	case token.AccessToken != "ACCESS_TOKEN":
		t.Errorf("tokens are not equal")
	case p.url != DefaultAuthURL:
		t.Errorf("expected url %s, got %s", DefaultAuthURL, p.url)
	case !bytes.Equal(p.body, []byte(expectedBody)):
		t.Errorf("expected body %s, got %s", p.body, p.url)
	}
}
