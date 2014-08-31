## Preface

Use this package to get a Go-based webapp off the ground in minutes.

## Features

* Go w/ Martini web stack
* Go templates
* Bower for frontend package management
* Heroku w/ New Relic integration

## Installation

### Common
1. Change instances of "goboilerplate" or "GOBOILERPLATE" to the name of your app, including the folder name
  * This last part is important for Go package discovery

### In development
1. `go get`
  * Only need to do this once
1. `npm install`
  * Only need to do this once
1. `bower install`
  * Do this whenever you change bower.json
1. `go run server.go`

### In production
1. `go get github.com/tools/godep`
  * Only need to do this once
1. `heroku config:add BUILDPACK_URL=https://github.com/ddollar/heroku-buildpack-multi.git`
  * Only need to do this once
1. `godep save` and commit changes
  * Do this whenever you add/remove a Go package
1. Push to Heroku repository
