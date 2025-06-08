package main

import (
	"nearby-web-app/presentation/http"

	"gofr.dev/pkg/gofr"
)

func main() {

	// initialise gofr object
	app := gofr.New()

	http.RegisterRoute(app)

	// Runs the server, it will listen on the default port 8000.
	// it can be over-ridden through configs
	app.Run()
}
