package push

import (
	"context"
	"time"

	"github.com/nburmi/huawei-push/push"
	"github.com/nburmi/huawei-push/token"
)

type Builder interface {
	SetAutoRefresherTokener(context.Context) Builder
	SetSubTimeTokener(time.Duration) Builder

	Build() (push.Pusher, error)
}

func New(p token.Params) Builder {
	return &builder{p: p}
}

type builder struct {
	p token.Params

	isRefresher bool
	subTime     time.Duration
	ctx         context.Context
}

func (b *builder) SetAutoRefresherTokener(ctx context.Context) Builder {
	b.isRefresher = true
	b.ctx = ctx
	return b
}

func (b *builder) SetSubTimeTokener(sub time.Duration) Builder {
	b.subTime = sub
	return b
}

func (b *builder) Build() (push.Pusher, error) {
	tok, err := token.New().SetByParams(b.p).Build()
	if err != nil {
		return nil, err
	}

	if b.isRefresher {
		tok, err = token.NewRefresher(b.ctx, tok).SetSubTime(b.subTime).Build()
	}

	pusher := push.New(b.p.ClientID, tok, b.p.HTTPDoer)

	return pusher, err
}
