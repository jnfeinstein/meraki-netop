package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
)

func main() {
	m := martini.Classic()

	m.Use(render.Renderer())

	m.Get("/login", func(req *http.Request, r render.Render) {
		r.HTML(200, "login", req.URL.Query())
	})

	m.Get("/clickthrough", func(req *http.Request, r render.Render) {
		r.HTML(200, "clickthrough", req.URL.Query())
	})

	m.Run()
}
