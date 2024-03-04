package home

import (
	"net/http"

	"github.com/book/help/internal/components/home"
	"github.com/book/help/internal/render"
)

func Index(w http.ResponseWriter, r *http.Request) {
	component := home.Home()

	render.Render(w, r, component)
}
