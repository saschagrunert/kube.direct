package main

import (
	"function"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/",
		func(res http.ResponseWriter, req *http.Request) {
			function.Handle(req.Context(), res, req)
		},
	)

	const (
		addr    = ":8081"
		timeout = 5 * time.Second
	)
	server := &http.Server{
		Addr:              addr,
		ReadHeaderTimeout: timeout,
	}

	log.Printf("Listening on %s", addr)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
