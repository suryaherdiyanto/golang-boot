package csrf

import (
	"crypto/sha1"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/sessions"
)

func New(key string) string {
	var secret = time.Now().UnixMicro()
	sha := sha1.New()

	sha.Write([]byte(fmt.Sprintf("%s__%s", strconv.Itoa(int(secret)), key)))
	hash := sha.Sum(nil)

	return fmt.Sprintf("%x", hash)
}

func Get(w http.ResponseWriter, r *http.Request) string {
	session := r.Context().Value("session").(sessions.Session)

	token, ok := session.Values["_csrf"]
	if !ok {
		token = New(os.Getenv("APP_KEY"))
		session.Values["_csrf"] = token
		session.Save(r, w)
	}

	return token.(string)
}
