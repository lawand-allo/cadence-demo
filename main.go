package main

import (
	delivery "cadence-demo/http"
	"cadence-demo/repository"
	"cadence-demo/usecase"
	"log"
	"net/http"
	"strconv"
	"time"
)

const PORT = 8080

func main() {
	orderRepository := repository.NewRepository()
	createOrderUsecase := usecase.NewCreateOrderUsecase(orderRepository)
	getOrderUsecase := usecase.NewGetOrderUsecase(orderRepository)
	handler := delivery.NewHandler(createOrderUsecase, getOrderUsecase)

	portString := strconv.Itoa(PORT)
	server := &http.Server{
		Addr:              ":" + portString,
		Handler:           handler,
		ReadHeaderTimeout: 60 * time.Second,
	}
	println("Starting server at localhost:" + portString)
	log.Fatal(server.ListenAndServe())
}
