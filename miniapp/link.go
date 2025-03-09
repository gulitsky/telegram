package miniapp

import (
	"fmt"
	"net/url"

	"github.com/gulitsky/telegram"
)

type Mode struct {
	slug string
}

type LinkOption func(url.Values)

var (
	Compact    = Mode{"compact"}
	Fullscreen = Mode{"fullscreen"}
)

func (m Mode) String() string {
	return m.slug
}

func WithStartApp(startApp string) LinkOption {
	return func(q url.Values) {
		if startApp != "" {
			q.Set("startapp", startApp)
		}
	}
}

func WithMode(mode Mode) LinkOption {
	return func(q url.Values) {
		if mode.slug != "" {
			q.Set("mode", mode.String())
		}
	}
}

func Link(botUsername, miniAppName string, options ...LinkOption) (string, error) {
	if miniAppName == "" {
		return "", fmt.Errorf("mini app name must not be empty")
	}
	if !shortNameRe.MatchString(miniAppName) {
		return "", fmt.Errorf("invalid mini app name: must contain 3-30 alphanumeric characters or underscores")
	}

	botLink, err := telegram.BotLink(botUsername)
	if err != nil {
		return "", fmt.Errorf("failed to generate bot link: %w", err)
	}

	u, err := url.Parse(botLink)
	if err != nil {
		return "", fmt.Errorf("failed to parse bot link: %w", err)
	}

	u = u.JoinPath(botLink, miniAppName)

	q := u.Query()
	for _, opt := range options {
		opt(q)
	}
	u.RawQuery = q.Encode()

	return u.String(), nil
}
