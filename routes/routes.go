package routes

import (
	"net/http"
	"strings"

	"github.com/wellingtonpires/api-stress-test/controller"
)

func PersonRoute(w http.ResponseWriter, r *http.Request) {
	if !strings.Contains(r.URL.Path, "/pessoas") {
		http.Error(w, "404 not found", http.StatusNotFound)
	}
	switch r.Method {
	case "GET":
		controller.GetPerson(w, r)
	case "POST":
		controller.Create(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func CountPersonRoute(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/contagem-pessoas" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}
	switch r.Method {
	case "GET":
		controller.GetCountPerson(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
