package sans_header_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var url string = "http://127.0.0.1:8000"

func RequestHeader(rw http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("content-type")
	fmt.Fprint(rw, contentType)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, url, nil)
	recorder := httptest.NewRecorder()

	request.Header.Add("Content-Type", "application/json")
	RequestHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}
