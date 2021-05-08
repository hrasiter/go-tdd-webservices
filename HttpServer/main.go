package main

import (
	"log"
	"net/http"

	"example.go.com/HttpServer/data"
)

func main() {
	ser := NewPlayerServer(data.NewInMemoryPlayerStore())
	log.Fatal(http.ListenAndServe(":5000", ser))
}
