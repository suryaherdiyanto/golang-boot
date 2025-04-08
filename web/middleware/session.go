package middleware

import (
	"context"
	"net/http"

	"github.com/gorilla/sessions"
)

func NewSession(store sessions.Store) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			s, err := store.Get(r, "kaspro_session")
			if err != nil {
				panic(err)
			}

			ctx := context.WithValue(r.Context(), "session", s)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetSession(r *http.Request) *sessions.Session {
	return r.Context().Value("session").(*sessions.Session)
}
