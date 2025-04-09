package response

import (
	"encoding/json"
	"net/http"
	"os"
)

type D map[string]interface{}

func JSON(w http.ResponseWriter, d interface{}) {
	w.Header().Add("Content-Type", "application/json")
	res, _ := json.Marshal(d)

	w.Write(res)
}

func TokenMismatched(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)

	JSON(w, D{
		"status":  http.StatusText(http.StatusBadRequest),
		"message": "CSRF Token Mismatched",
	})
}

func InternalServerError(w http.ResponseWriter, message string, err error) {
	w.WriteHeader(http.StatusInternalServerError)

	d := D{
		"status":  http.StatusText(http.StatusInternalServerError),
		"message": message,
	}

	if os.Getenv("APP_ENV") == "production" {
		d["exeption"] = err.Error()
	}

	JSON(w, d)
}
