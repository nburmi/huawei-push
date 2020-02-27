package push

import "fmt"

const success = "80000000"

type PusherError struct {
	code string
	desc string
}

func (p *PusherError) Error() string {
	return fmt.Sprintf("error code: %s, description: %s", p.code, p.desc)
}

// NewCheckPusher errors check error and return new error
func NewCheckPusher(p Pusher) Pusher {
	return &checker{p}
}

type checker struct {
	Pusher
}

func (p *checker) Push(m *Message) (*Response, error) {
	resp, err := p.Pusher.Push(m)
	if err != nil {
		return resp, err
	}

	return resp, p.checkResponse(resp)
}

func (p *checker) PushValidate(m *Message) (*Response, error) {
	resp, err := p.Pusher.PushValidate(m)
	if err != nil {
		return resp, err
	}

	return resp, p.checkResponse(resp)
}

func (p *checker) checkResponse(r *Response) error {
	if r.Code != success {
		return &PusherError{code: r.Code, desc: r.Message}
	}

	return nil
}
