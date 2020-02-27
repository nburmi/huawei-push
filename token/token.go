package token

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/nburmi/huawei-push/common"
)

const (
	DeafultGrantType = "client_credentials"
	DefaultAuthURL   = "https://oauth-login.cloud.huawei.com/oauth2/v2/token"
)

/*
Params for gettting access token.
	- ClientID:  APP ID obtained when you create the app on HUAWEI Developer
	- ClientSecret: App secret key obtained when you create the app on HUAWEI Developer
	- GrantType: if empty then value will be DeafultGrantType
	- URL: if empty then value will be DefaultApiTokenURL
	- HttpPoster: interface Post(url, contentType string, body io.Reader) (resp *http.Response, err error)
*/
type Params struct {
	ClientID     string
	ClientSecret string
	GrantType    string
	URL          string
	common.HTTPDoer
}

//Client Password Mode
type Token struct {
	StatusCode       int
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	TokenType        string `json:"token_type"`
	Error            int    `json:"error"`
	SubError         int    `json:"sub_error"`
	ErrorDescription string `json:"error_description"`
}

type Tokener interface {
	Get() (*Token, error)
}

type BuilderTokener interface {
	SetByParams(Params) BuilderTokener

	SetID(string) BuilderTokener
	SetSecret(string) BuilderTokener
	SetHTTPDoer(common.HTTPDoer) BuilderTokener
	SetGrantType(string) BuilderTokener

	Build() (Tokener, error)
}

func New() BuilderTokener {
	return &tokenerReal{}
}

type tokenerReal struct {
	Params
}

func (t *tokenerReal) SetByParams(p Params) BuilderTokener {
	t.Params = p
	return t
}

func (t *tokenerReal) SetID(id string) BuilderTokener {
	t.ClientID = id
	return t
}

func (t *tokenerReal) SetSecret(s string) BuilderTokener {
	t.ClientSecret = s
	return t
}

func (t *tokenerReal) SetHTTPDoer(p common.HTTPDoer) BuilderTokener {
	t.HTTPDoer = p
	return t
}

func (t *tokenerReal) SetGrantType(g string) BuilderTokener {
	t.GrantType = g
	return t
}

func (t *tokenerReal) Get() (*Token, error) {
	vals := url.Values{
		"grant_type":    []string{t.GrantType},
		"client_secret": []string{t.ClientSecret},
		"client_id":     []string{t.ClientID},
	}

	req, err := http.NewRequest("POST", t.URL, strings.NewReader(vals.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := t.Do(req)
	if err != nil {
		return &Token{StatusCode: resp.StatusCode}, err
	}

	var tok Token
	err = json.NewDecoder(resp.Body).Decode(&tok)
	if err != nil && !errors.Is(err, io.EOF) {
		return nil, err
	}

	resp.Body.Close()

	return &tok, nil
}

// Build and validate params
func (t *tokenerReal) Build() (Tokener, error) {
	var err error

	switch {
	case t.Params.ClientID == "":
		err = errors.New("ClientID is empty")
	case t.Params.ClientSecret == "":
		err = errors.New("ClientSecret is empty")
	case t.Params.HTTPDoer == nil:
		err = errors.New("HTTPDoer is not set")
	case t.Params.GrantType == "":
		t.Params.GrantType = DeafultGrantType
		fallthrough
	case t.Params.URL == "":
		t.Params.URL = DefaultAuthURL
	}

	return t, err
}
