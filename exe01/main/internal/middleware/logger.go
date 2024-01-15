package middleware

import (
	"fmt"
	"net/http"
	"time"
)

type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		fmt.Println(">>> Verb: ", r.Method, "Path: ", r.URL.Path)
		fmt.Println(">>> Time: ", time.Now().Format(time.RFC3339))
		fmt.Println(">>> Size in bytes: ", r.ContentLength)

		next.ServeHTTP(w, r)
	})
}
