package response

import "net/http"

type D map[string]interface{}

func TokenMismatched(w http.ResponseWriter) {
	w.WriteHeader(400)
	w.Write([]byte("CSRF Token Mismatched"))
}
