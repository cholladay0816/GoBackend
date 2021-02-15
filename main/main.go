package main

import (
	"GoBackend/routes"
	"log"
	"net/http"
)

func main() {
	routes.Setup()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
