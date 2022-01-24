package sans_http_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func handlerHelloWorld(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello World!")
}

var target string = "http://127.0.0.1:8000"

func TestHandlerHelloWorld(t *testing.T) {
	req := httptest.NewRequest("GET", target, nil) // new http
	rec := httptest.NewRecorder()

	handlerHelloWorld(rec, req)

	res := rec.Result()
	body, _ := io.ReadAll(res.Body)
	bodyString := string(body)

	fmt.Println(bodyString)

}
