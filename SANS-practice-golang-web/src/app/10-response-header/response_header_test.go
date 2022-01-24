package sans_header_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ReponseHeader(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("X-CREATE-BY", "SANS")
}

var url string = "http://127.0.0.1:8000"

func TestReponseHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, url, nil)
	recorder := httptest.NewRecorder()
	ReponseHeader(recorder, request)

	createBy := recorder.Header().Get("x-create-by")

	fmt.Println(createBy)
}
