package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/sessions"
	"github.com/petaki/inertia-go"
	"github.com/suryaherdiyanto/kaspro/web/handler"
	webmiddleware "github.com/suryaherdiyanto/kaspro/web/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	mux := chi.NewRouter()
	server := http.Server{
		Addr:    ":" + os.Getenv("APP_PORT"),
		Handler: mux,
	}
	store := sessions.NewCookieStore([]byte(os.Getenv("APP_KEY")))
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	fs := http.FileServer(http.Dir("./public"))
	jsResources := http.FileServer(http.Dir("./resources/js/"))

	store.Options.MaxAge = 3600
	store.Options.HttpOnly = true
	store.Options.SameSite = http.SameSiteLaxMode
	store.Options.Secure = os.Getenv("APP_ENV") == "production"

	inertiaManager := inertia.New("http://localhost:3000", "./resources/tmpl/app.go.html", os.Getenv("INERTIA_VERSION"))

	handler := handler.New(inertiaManager, db)
	mux.Use(middleware.Logger, webmiddleware.NewSession(store), webmiddleware.ValidateCSRF, inertiaManager.Middleware)

	mux.Handle("/asset/*", http.StripPrefix("/asset/", fs))
	mux.Handle("/js/*", http.StripPrefix("/js/", jsResources))

	mux.Get("/", handler.Home)

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
