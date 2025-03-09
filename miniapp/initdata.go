package miniapp

import (
	"context"
	"encoding/base64"
	"errors"
	"net/http"
	"strings"
	"time"

	initdata "github.com/telegram-mini-apps/init-data-golang"
)

type contextKey struct {
	name string
}

const bearerPrefix string = "Bearer "

var (
	initDataKey = contextKey{"initData"}
)

func Auth(
	botToken string,
	expTime time.Duration,
) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				writeError(w, http.StatusUnauthorized, "missing Authorization header")
				return
			}

			if !strings.HasPrefix(authHeader, bearerPrefix) {
				writeError(w, http.StatusBadRequest, "invalid Authorization header: Bearer prefix required")
				return
			}

			token := strings.TrimPrefix(authHeader, bearerPrefix)
			if token == "" {
				writeError(w, http.StatusBadRequest, "empty token")
				return
			}

			rawInitData, err := base64.URLEncoding.DecodeString(token)
			if err != nil {
				writeError(w, http.StatusBadRequest, "invalid base64 token")
				return
			}

			if err := initdata.Validate(string(rawInitData), botToken, expTime); err != nil {
				switch {
				case errors.Is(err, initdata.ErrUnexpectedFormat):
					writeError(w, http.StatusBadRequest, "invalid init data format")
					return
				case errors.Is(err, initdata.ErrSignMissing):
					writeError(w, http.StatusBadRequest, "missing hash parameter in init data")
					return
				case errors.Is(err, initdata.ErrAuthDateMissing):
					writeError(w, http.StatusBadRequest, "missing or invalid auth_date parameter in init data")
					return
				case errors.Is(err, initdata.ErrExpired):
					writeError(w, http.StatusUnauthorized, "init data has expired")
					return
				case errors.Is(err, initdata.ErrSignInvalid):
					writeError(w, http.StatusUnauthorized, "invalid init data signature")
					return
				default:
					writeError(w, http.StatusUnauthorized, "init data validation failed: "+err.Error())
					return
				}
			}

			initData, err := initdata.Parse(string(rawInitData))
			if err != nil {
				writeError(
					w,
					http.StatusBadRequest,
					"failed to parse init data",
				)
				return
			}

			ctx := context.WithValue(r.Context(), initDataKey, &initData)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func InitDataFromContext(ctx context.Context) *initdata.InitData {
	if initData, ok := ctx.Value(initDataKey).(*initdata.InitData); ok {
		return initData
	}
	return nil
}

func writeError(w http.ResponseWriter, status int, msg string) {
	http.Error(w, msg, status)
}
