package token

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

type mockTokener struct {
	count     int64
	expiresIN int
}

func (m *mockTokener) Get() (*Token, error) {
	t := &Token{
		StatusCode:  http.StatusAccepted,
		AccessToken: strconv.FormatInt(m.count, 10),
		ExpiresIn:   1,
	}
	atomic.AddInt64(&m.count, 1)
	return t, nil
}

// checking when token is refreshing
func TestRefresherCheckWithExpire(t *testing.T) {

	c, cancel := context.WithCancel(context.Background())
	tok := &mockTokener{expiresIN: 1}
	refresher, _ := NewRefresher(c, tok).Build()

	for i := 0; i < 5; i++ {
		tk, _ := refresher.Get()
		expected := strconv.Itoa(i)

		if !strings.EqualFold(tk.AccessToken, expected) {
			t.Errorf("wrong access token, expected %s, got %s", expected, tk.AccessToken)
		}

		time.Sleep(time.Second + time.Millisecond*100)
	}

	cancel()

	time.Sleep(time.Second + time.Millisecond*100)
	tk, _ := refresher.Get()
	expected := "5"
	if !strings.EqualFold(tk.AccessToken, expected) {
		t.Errorf("wrong access token, expected %s, got %s", expected, tk.AccessToken)
	}
}

func TestRefresherCheckWithCustomExpire(t *testing.T) {

	c, cancel := context.WithCancel(context.Background())
	tok := &mockTokener{expiresIN: 1}
	refresher, _ := NewRefresher(c, tok).CustomExpireTime(100 * time.Millisecond).Build()

	for i := 0; i < 5; i++ {
		tk, _ := refresher.Get()
		expected := strconv.Itoa(i)

		if !strings.EqualFold(tk.AccessToken, expected) {
			t.Errorf("wrong access token, expected %s, got %s", expected, tk.AccessToken)
		}

		time.Sleep(time.Millisecond * 110)
	}

	cancel()

	time.Sleep(time.Second)
	tk, _ := refresher.Get()
	expected := "5"
	if !strings.EqualFold(tk.AccessToken, expected) {
		t.Errorf("wrong access token, expected %s, got %s", expected, tk.AccessToken)
	}
}
