package main

import (
	"go-postgres/router"
	"net/http"
)

func main() {

    app := &router.App{
		Addr: ":9090",
	}

	http.ListenAndServe(app.Addr , app.Router())
}