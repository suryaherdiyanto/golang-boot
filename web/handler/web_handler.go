package handler

import (
	"net/http"

	"github.com/suryaherdiyanto/kaspro/web/response"
)

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	h.Inertia.Render(w, r, "Index", response.D{
		"foo": "bar",
	})
}
