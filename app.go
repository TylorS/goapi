package goapi

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tylors/goapi/middleware"
	"github.com/tylors/goapi/router"
)

type App struct {
	DB          *sql.DB
	Router      *router.Router
	Middlewares []middleware.Middleware
}

func (a *App) Handle(path string, handler http.Handler) *mux.Route {
	return a.Router.Handle(path, handler)
}

func (a *App) HandleFunc(path string, handler http.HandlerFunc) *mux.Route {
	return a.Router.Handle(path, handler)
}

func (a *App) Use(middlewares ...middleware.Middleware) {
	for _, middleware := range middlewares {
		a.Router.Middlewares = append(a.Middlewares, middleware)
	}
}

func (a *App) RegisterBlueprint(blueprint BluePrint, prefix string) {
	blueprint.RegisterRouter(a, prefix)
	blueprint.Routes()
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.Router.ServeHTTP(w, r)
}

func (a *App) Run(addr string) {
	fmt.Println("Running server on:", addr)
	log.Fatal(http.ListenAndServe(addr, a))
}

func New() *App {
	app := &App{
		Router: router.New(),
	}
	app.Use(middleware.Logger, middleware.Recovery)
	return app
}
