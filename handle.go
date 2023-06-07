package function

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Handle an HTTP Request.
func Handle(_ context.Context, res http.ResponseWriter, req *http.Request) {
	/*
	 * YOUR CODE HERE
	 *
	 * Try running `go test`.  Add more test as you code in `handle_test.go`.
	 */

	fmt.Println("Received request")
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
