// +build !heroku

package config

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

const HEROKU = false

func Initialize(m *martini.ClassicMartini) {
	fmt.Println("Running in debug environment")

	m.Use(martini.Static("private"))

	m.Use(render.Renderer(render.Options{
		Layout: "app-devel",
	}))
}
