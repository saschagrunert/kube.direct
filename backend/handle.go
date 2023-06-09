package function

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"function/internal/server"
)

var grpcServerDead atomic.Bool

func init() {
	srv, err := server.New()
	if err != nil {
		log.Fatalf("Unable to init gRPC server: %v", err)
	}

	go func() {
		if err := srv.Serve(); err != nil {
			log.Printf("Unable to serve gRPC: %v", err)
			grpcServerDead.Store(true)
		}
	}()
}

// Handle an HTTP Request.
func Handle(_ context.Context, res http.ResponseWriter, req *http.Request) {
	if grpcServerDead.Load() {
		log.Fatal("The gRPC server is not running any more, exiting")
	}

	log.Print("Received request")
	prettyPrint(req, os.Stdout) // echo to local output
	prettyPrint(req, res)
}

func prettyPrint(req *http.Request, w io.Writer) {
	fmt.Fprintf(w, "%v %v %v %v\n", req.Method, req.URL, req.Proto, req.Host)
	for k, vv := range req.Header {
		for _, v := range vv {
			fmt.Fprintf(w, "  %v: %v\n", k, v)
		}
	}
}
