package app

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(w http.ResponseWriter, r *http.Request) {

	qpName := r.URL.Query().Get("name")

	//TODO  SET COOKIE
	cookie := new(http.Cookie)
	cookie.Name = "X-NAME"
	cookie.Value = qpName
	cookie.Path = "/"

	http.SetCookie(w, cookie)
	fmt.Fprint(w, "OKE")

}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("X-NAME")
	if cookie.Name != "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "NOT COOKIE SET")
	} else {
		name := cookie.Value
		fmt.Fprintf(w, "Hello %s", name)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: mux,
	}

	server.ListenAndServe()
}

func TestSetCookie(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "http://localhost:8000/?name=Sanrian", nil)
	rec := httptest.NewRecorder()

	SetCookie(rec, req)

	cookies := rec.Result().Cookies()

	for _, cookie := range cookies {
		fmt.Printf("Cookie %s:%s \n", cookie.Name, cookie.Value)
	}

}

func TestGetCookie(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8000", nil) // no query params
	// req := httptest.NewRequest(http.MethodGet, "http://localhost:8000?name=sandrian", nil) // with query params
	rec := httptest.NewRecorder()

	qpName := req.URL.Query().Get("name")

	if qpName == "" {
		fmt.Println("NO FOUND FOR COOKIE")
	}

	cookie := new(http.Cookie)
	cookie.Name = "X-NAME"
	cookie.Value = qpName

	req.AddCookie(cookie)
	GetCookie(rec, req)

	response := rec.Result()
	status := response.Status
	statusCode := response.StatusCode
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(string(bodyString))
	fmt.Println(status)
	fmt.Println(statusCode)

}
