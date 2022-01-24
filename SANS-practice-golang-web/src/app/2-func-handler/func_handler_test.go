package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func TestHandlerFunc(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8000",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic("ERROR")
	}

}
