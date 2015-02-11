package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tylors/goapi/middleware"
)

type Router struct {
	Mux         *mux.Router
	Middlewares []middleware.Middleware
}

func (r *Router) Handle(path string, handler http.Handler) *mux.Route {
	return r.Mux.Handle(path, middleware.New(r.Middlewares...).Then(handler))
}

func (r *Router) HandleFunc(path string, handler http.HandlerFunc) *mux.Route {
	return r.Mux.Handle(path, middleware.New(r.Middlewares...).ThenFunc(handler))
}

func (r *Router) Subrouter(prefix string) *Router {
	return &Router{
		Mux: r.Mux.PathPrefix(prefix).Subrouter().StrictSlash(true),
	}
}

func New() *Router {
	return &Router{
		Mux: mux.NewRouter().StrictSlash(true),
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.Mux.ServeHTTP(w, req)
}
