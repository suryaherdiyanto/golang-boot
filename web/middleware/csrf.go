package middleware

import (
	"net/http"
	"slices"

	"github.com/suryaherdiyanto/kaspro/web/response"
)

func ValidateCSRF(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var mustValidateMethods = []string{"POST", "PUT", "DELETE", "PATCH"}

		if slices.Index(mustValidateMethods, r.Method) > -1 {
			tokenHeader := r.Header.Get("X-CSRF")

			if tokenHeader == "" {
				response.TokenMismatched(w)
				return
			}

			session := GetSession(r)

			tokenSession, ok := session.Values["_csrf"]

			if !ok {
				response.TokenMismatched(w)
				return
			}

			if tokenHeader != tokenSession {
				response.TokenMismatched(w)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
