package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/sessions"
	"github.com/petaki/inertia-go"
	webmiddleware "github.com/suryaherdiyanto/kaspro/web/middleware"
	"github.com/suryaherdiyanto/kaspro/web/response"
)

func main() {
	mux := chi.NewRouter()
	server := http.Server{
		Addr:    ":" + os.Getenv("APP_PORT"),
		Handler: mux,
	}
	store := sessions.NewCookieStore([]byte(os.Getenv("APP_KEY")))

	fs := http.FileServer(http.Dir("./public"))
	jsResources := http.FileServer(http.Dir("./resources/js/"))

	store.Options.MaxAge = 3600
	store.Options.HttpOnly = true
	store.Options.SameSite = http.SameSiteLaxMode
	store.Options.Secure = os.Getenv("APP_ENV") == "production"

	inertiaManager := inertia.New("http://localhost:3000", "./resources/tmpl/app.go.html", os.Getenv("INERTIA_VERSION"))

	mux.Use(middleware.Logger)
	mux.Use(webmiddleware.NewSession(store))
	mux.Use(webmiddleware.ValidateCSRF)
	mux.Use(inertiaManager.Middleware)

	mux.Handle("/asset/*", http.StripPrefix("/asset/", fs))
	mux.Handle("/js/*", http.StripPrefix("/js/", jsResources))

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		err := inertiaManager.Render(w, r, "Index", response.D{
			"foo": "bar",
		})

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
		}
	})

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
