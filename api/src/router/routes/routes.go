package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route representa uma rota da API
type Route struct {
	Uri                   string
	Method                string
	Function              func(w http.ResponseWriter, r *http.Request)
	RequireAuthentication bool
}

func Configure(r *mux.Router) *mux.Router {
	routes := UserRoutes

	for _, route := range routes {
		r.HandleFunc(route.Uri, route.Function).Methods(route.Method)
	}

	return r
}
