package goapi

import "github.com/tylors/goapi/middleware"

type BluePrint interface {
	RegisterRouter(*App, string)
	Initialize()
	Routes()
}

type Blueprint struct {
	App
}

func (b *Blueprint) RegisterRouter(app *App, prefix string) {
	b.Router = app.Router.Subrouter(prefix)
}

func (b *Blueprint) Initialize() {
	b.Use(middleware.Logger, middleware.Recovery)
}
