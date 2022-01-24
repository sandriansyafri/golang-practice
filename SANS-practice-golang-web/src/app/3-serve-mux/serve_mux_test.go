package serve_mux_test

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServeMux(t *testing.T) {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hello World!")
	})
	mux.HandleFunc("/rian", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hello World Rian!")
	})

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic("ERROR !")
	}

}
