package utils

import (
	"fmt"
	"net/http"
)

type App struct {
	server http.Server
	routes *(map[string]map[string]http.HandlerFunc)
}

func (app *App) Create() {
	app.server = http.Server{}
	app.routes = &map[string]map[string]http.HandlerFunc{}
	fmt.Println("App created")
}

func (app *App) Start(addr string) {
	app.server.Addr = addr
	app.server.Handler = app
	app.server.ListenAndServe()
	fmt.Println("Goapi Started")

}

func (app *App) Register(method string, endpoint string, h http.HandlerFunc) {
	routes := *(app.routes)
	if _, ok := routes[method]; ok {
		routes[method][endpoint] = h
	} else {
		routes[method] = map[string]http.HandlerFunc{}
		fmt.Println(app.routes)
		routes[method][endpoint] = h
	}
	fmt.Println("add", method, endpoint)
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	endpoint := r.URL.Path
	routes := *(app.routes)

	if h, ok := routes[method][endpoint]; ok {
		h(w, r)
	}
}
