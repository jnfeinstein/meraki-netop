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
1. `gem install compass` to install the compass stylesheet processor
  * You may (probably) need to `sudo` this command
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


## Folder structure

### config

By default this folder contains devel.go and heroku.go, which both define the package ```config```.  Only one of them will be imported by server.go depending on whether the build tag ```heroku``` is defined (production.go) or not (development.go).  You can define environment specific code in these files.  New Relic is initialized in heroku.go, which is only used in the Heroku environment.

### Godeps

This folder is auto-generated by the godep binary when you run ```godep save```.  It is used to store Go dependencies, which Heroku requires.

### node_modules

This folder is produces by nodeJS and contains nodeJS packages.  This project uses ```bower``` by default, which is installed to this directory using ```npm install```.

### public

Any static assets should be places in this folder.  The server knows how to server anything in this folder without a specific handler in server.go.  The ```vendor``` folder is produced by ```bower install``` and contains any bower-managed packages.

### templates

Go HTML templates should be placed in this folder.

### Everything else

You're on your own!  You can set up your Go code however you wish.  This project uses server.go as its main file, but you can modify this however you wish.
