package main

import (
	delivery "cadence-demo/http"
	"log"
	"strconv"
)

const PORT = 8080

func main() {
	router := delivery.NewRouter()
	server := delivery.NewServer(router, PORT)
	portString := strconv.Itoa(PORT)
	println("Starting server at localhost:" + portString)
	log.Fatal(server.ListenAndServe())
}
