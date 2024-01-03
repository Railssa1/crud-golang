package main

import (
	"loja/routes"
	"net/http"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":3000", nil)
}
