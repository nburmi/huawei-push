package token

import (
	"context"
	"errors"
	"sync"
	"time"
)

const defaultTimerDuration = time.Second

// tokenRefresher - store active Token and refresh then token is expired
type tokenRefresher struct {
	token *Token
	err   error

	once  *sync.Once
	mu    *sync.RWMutex
	timer *time.Timer
	ctx   context.Context

	sub    time.Duration
	custom time.Duration
	Tokener
}

type RefresherBuilder interface {
	Build() (Tokener, error)
	SetSubTime(sub time.Duration) RefresherBuilder
	CustomExpireTime(t time.Duration) RefresherBuilder
}

// context using for stoping timer refresher
func NewRefresher(c context.Context, t Tokener) RefresherBuilder {
	if c == nil {
		c = context.Background()
	}

	return &tokenRefresher{Tokener: t, ctx: c, once: &sync.Once{}, mu: &sync.RWMutex{}}
}

// SubTime - update the token more often for sub time
func (t *tokenRefresher) SetSubTime(sub time.Duration) RefresherBuilder {
	t.sub = sub
	return t
}

// SetCustomExpireTime - update the token more often for sub time
func (t *tokenRefresher) CustomExpireTime(c time.Duration) RefresherBuilder {
	t.custom = c
	return t
}

func (t *tokenRefresher) Build() (Tokener, error) {
	var err error

	switch {
	case t.Tokener == nil:
		err = errors.New("Tokener is nil")
	case t.sub > 0 && t.custom > 0 && t.sub >= t.custom:
		err = errors.New("sub must be less than custom")
	}

	return t, err
}

// initAndStart - get Token and Start Timer
func (t *tokenRefresher) initAndStart() {
	t.refreshToken()

	t.timer = time.NewTimer(t.getExpiteTime())
	go t.eventLoop()
}

func (t *tokenRefresher) Get() (*Token, error) {
	t.once.Do(t.initAndStart)

	t.mu.RLock()
	defer t.mu.RUnlock()

	return t.token, t.err
}

func (t *tokenRefresher) eventLoop() {
	for {
		select {
		case <-t.timer.C:
			t.refreshToken()
			t.timer.Reset(t.getExpiteTime())
		case <-t.ctx.Done():
			t.timer.Stop()
		}
	}
}

func (t *tokenRefresher) refreshToken() {
	token, err := t.Tokener.Get()

	t.mu.Lock()
	t.token, t.err = token, err
	t.mu.Unlock()
}

func (t *tokenRefresher) getExpiteTime() time.Duration {
	if t.token == nil {
		return defaultTimerDuration
	}

	resetTime := time.Second * time.Duration(t.token.ExpiresIn)
	if t.custom > 0 {
		resetTime = t.custom
	}

	if tm := resetTime - t.sub; tm > 0 {
		resetTime = tm
	}

	return resetTime
}
