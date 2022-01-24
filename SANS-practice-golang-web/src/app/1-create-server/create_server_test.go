package create_server_test

import (
	"net/http"
	"testing"
)

func TestGolangWeb(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8000",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic("ERROR")
	}

}
