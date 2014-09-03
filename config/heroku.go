// +build heroku

package config

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/gorelic"
	"github.com/martini-contrib/render"
	"os"
)

const HEROKU = true
const APPNAME = "GOBOILERPLATE"

func Initialize(m *martini.ClassicMartini) {
	fmt.Println("Running in production environment")

	m.Use(render.Renderer(render.Options{
		Layout: "app-heroku",
	}))

	NEW_RELIC_LICENSE_KEY := os.Getenv("NEW_RELIC_LICENSE_KEY")
	if len(NEW_RELIC_LICENSE_KEY) > 0 {
		gorelic.InitNewrelicAgent(NEW_RELIC_LICENSE_KEY, APPNAME, true)
		m.Use(gorelic.Handler)
	}
}
