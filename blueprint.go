package goapi

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
