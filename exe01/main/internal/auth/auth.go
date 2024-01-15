package auth

import "errors"

var (
	ErrAuthTokenInternal = errors.New("auth token internal error")
)

type AuthToken interface {
	Auth(token string) (err error)
}
