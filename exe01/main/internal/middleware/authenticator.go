package middleware

import "net/http"

type Authenticator struct {
	token string
}

func NewAuthenticator(token string) *Authenticator {
	return &Authenticator{
		token: token,
	}
}

func (a *Authenticator) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != a.token {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
