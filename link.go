package telegram

import (
	"fmt"
	"net/url"
)

type LinkOption func(url.Values)

const (
	scheme = "https"
	host   = "t.me"
)

func WithStart(start string) LinkOption {
	return func(q url.Values) {
		if start != "" {
			q.Set("start", start)
		}
	}
}

func BotLink(username string, options ...LinkOption) (string, error) {
	if username == "" {
		return "", fmt.Errorf("username must not be empty")
	}
	if !usernameRe.MatchString(username) {
		return "", fmt.Errorf("invalid username: must contain 5-32 alphanumeric characters or underscores")
	}

	u := url.URL{
		Scheme: scheme,
		Host:   host,
		Path:   username,
	}

	q := u.Query()
	for _, opt := range options {
		opt(q)
	}
	u.RawQuery = q.Encode()

	return u.String(), nil
}
