package belajar_golang_web

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "X-PCK"
	cookie.Value = r.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(w, cookie)
	fmt.Fprint(w, "success create cookie")
}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("X-PCK")
	if err != nil {
		fmt.Fprint(w, "No Cookie")
	} else {
		fmt.Fprint(w, cookie.Value)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

func TestSetCookie(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/set-cookie?name=andi", nil)
	res := httptest.NewRecorder()

	SetCookie(res, req)

	cookies := res.Result().Cookies()
	for _, cookie := range cookies {
		fmt.Printf("Cookie %s:%s\n", cookie.Name, cookie.Value)
	}
}

func TestGetCookie(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/get-cookie", nil)
	cookie := new(http.Cookie)
	cookie.Name = "X-PCK"
	cookie.Value = "andi"
	req.AddCookie(cookie)

	res := httptest.NewRecorder()

	GetCookie(res, req)
	fmt.Println(res.Body.String())
}
