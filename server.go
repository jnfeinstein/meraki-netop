package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
)

func main() {
	m := martini.Classic()

	m.Use(render.Renderer())

	m.Get("/", func(req *http.Request, r render.Render) {
		r.HTML(200, "app", req.URL.Query())
	})

	m.Run()
}
