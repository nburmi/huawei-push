package token

import "fmt"

// NewResponseChecker token errors check error and return new error
func NewResponseChecker(t Tokener) Tokener {
	return &checker{t}
}

type TokenerError struct {
	code int
	sub  int
	desc string
}

func (t *TokenerError) Error() string {
	return fmt.Sprintf("error code: %d, sub error code:%d, description: %s", t.code, t.sub, t.desc)
}

type checker struct {
	Tokener
}

func (t *checker) Get() (*Token, error) {
	tok, err := t.Tokener.Get()
	if err != nil {
		return tok, err
	}

	if tok.Error != 0 || tok.SubError != 0 || tok.ErrorDescription != "" {
		err = &TokenerError{code: tok.Error, sub: tok.SubError, desc: tok.ErrorDescription}
	}

	return tok, err
}
