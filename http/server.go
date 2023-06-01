package http

import (
	"net/http"
	"strconv"
	"time"
)

func NewServer(router http.Handler, port int) *http.Server {
	server := &http.Server{
		Addr:              ":" + strconv.Itoa(port),
		Handler:           router,
		ReadHeaderTimeout: 60 * time.Second,
	}
	return server
}
