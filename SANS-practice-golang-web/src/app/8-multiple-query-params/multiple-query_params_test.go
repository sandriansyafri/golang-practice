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
func handlerMultipleQueryParams(rw http.ResponseWriter, r *http.Request) {
	firstname := r.URL.Query().Get("firstname")
	lastname := r.URL.Query().Get("lastname")

	fmt.Fprintf(rw, "Hello %s %s", firstname, lastname)

}

var target1 string = "http://127.0.0.1:8000?name=asdfasdfsadf"
var target2 string = "http://127.0.0.1:8000?firstname=Sandrian&lastname=Syafri"

func TestHandlerQueryParams(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, target1, nil)
	rec := httptest.NewRecorder()

	handlerQueryParams(rec, req)

	res := rec.Result()
	body, _ := io.ReadAll(res.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}
func TestHandlerMultipleQueryParams(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, target2, nil)
	rec := httptest.NewRecorder()

	handlerMultipleQueryParams(rec, req)

	res := rec.Result()
	body, _ := io.ReadAll(res.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}
