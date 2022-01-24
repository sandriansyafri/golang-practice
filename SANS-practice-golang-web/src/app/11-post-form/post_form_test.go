package post_form_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var url string = "http://127.0.0.1:8000"

func FormPost(w http.ResponseWriter, r *http.Request) {
	// TODO MAGIC
	firstname := r.PostFormValue("firstname")
	lastname := r.PostFormValue("lastname")

	// TODO MANUAL
	// err := r.ParseForm()
	// if err != nil {
	// 	panic("404")
	// }
	// firstname := r.PostForm.Get("firstname")
	// lastname := r.PostForm.Get("lastname")

	fmt.Fprintf(w, "Hello %s %s", firstname, lastname)
}

func TestFormPost(t *testing.T) {
	reqbody := strings.NewReader("firstname=Sandrian&lastname=Syafri")

	req := httptest.NewRequest("POST", url, reqbody)
	rec := httptest.NewRecorder()

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	FormPost(rec, req)

	response := rec.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)

}
