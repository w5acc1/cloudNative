package main

import (
	"io"
	"net/http"
)

func healthz(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			w.Header().Add(name, h)
		}
	}
	io.WriteString(w, "this is healthz")

	w.WriteHeader(http.StatusOK)
}

func main() {

	http.HandleFunc("/healthz", healthz)

	http.ListenAndServe(":8090", nil)

}
