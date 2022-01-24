package app

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "INTERNAL ERROR")
	} else {
		fmt.Fprintf(w, "HELLO %s", name)
	}
}

func TestResponseCode(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "http://localhost:8000/", nil)
	rec := httptest.NewRecorder()

	ResponseCode(rec, req)

	response := rec.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	status := response.Status
	statusCode := response.StatusCode

	fmt.Println("status : ", status)
	fmt.Println("statusCode : ", statusCode)
	fmt.Println(bodyString)

}
