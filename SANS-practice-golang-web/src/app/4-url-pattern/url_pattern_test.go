package url_pattern

import (
	"fmt"
	"net/http"
	"testing"
)

func TestUrlPattern(t *testing.T) {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hello World!")
	})
	mux.HandleFunc("/images/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "images/")
	})
	mux.HandleFunc("/images/rian/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "images/rian/")
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
