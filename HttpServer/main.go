package main

import (
	"log"
	"net/http"

	"example.go.com/HttpServer/data"
	"example.go.com/HttpServer/handler"
)

func main() {
	ser := handler.NewPlayerServer(data.NewInMemoryPlayerStore())
	log.Fatal(http.ListenAndServe(":5000", ser))
}
