package goapi

import "database/sql"

type BluePrint interface {
	RegisterRouter(*App, string)
	Initialize()
	Routes()
}

type Blueprint struct {
	App
	DB *sql.DB
}

func (b *Blueprint) RegisterRouter(app *App, prefix string) {
	b.DB = app.DB
	b.Router = app.Router.Subrouter(prefix)
}
