package handlers

import (
	"net/http"

	"github.com/a-h/templ"

	"github.com/david-daxzbr/english-app/views/page"
)

func Home(w http.ResponseWriter, r *http.Request) {
	templ.Handler(page.Home("home")).ServeHTTP(w, r)
}
