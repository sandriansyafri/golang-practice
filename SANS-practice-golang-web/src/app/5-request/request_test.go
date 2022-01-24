package request_test

import (
	"fmt"
	"net/http"
	"testing"
)

var handler http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, r.Method)
	fmt.Fprint(rw, r.RequestURI)
}

func TestRequest(t *testing.T) {

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic("404")
	}

}
