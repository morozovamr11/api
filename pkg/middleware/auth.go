package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

func IsAuthed(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authed := r.Header.Get("Authorization")
		token := strings.TrimPrefix(authed, "Bearer ")
		fmt.Println(token)
		next.ServeHTTP(w, r)
	})
}
