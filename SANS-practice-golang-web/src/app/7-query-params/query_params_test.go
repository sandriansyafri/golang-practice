package sans_query_params

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func handlerQueryParams(rw http.ResponseWriter, r *http.Request) {
	queryName := r.URL.Query().Get("name")
	if queryName == "" {
		fmt.Fprint(rw, "Hello")
	} else {
		fmt.Fprintf(rw, "Hello %s", queryName)
	}
}

var target string = "http://127.0.0.1:8000?name=asdfasdfsadf"

func TestHandlerQueryParams(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, target, nil)
	rec := httptest.NewRecorder()

	handlerQueryParams(rec, req)

	res := rec.Result()
	body, _ := io.ReadAll(res.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}
