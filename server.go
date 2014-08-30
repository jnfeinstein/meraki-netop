package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"goboilerplate/config"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	m := martini.Classic()

	config.Initialize(m)

	m.Use(render.Renderer(render.Options{
		Layout: "application",
	}))

	m.Get("/", func(r render.Render) {
		r.HTML(200, "index", nil)
	})

	m.Run()
}
