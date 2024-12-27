package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(writter http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(writter, "Hello")
	} else {
		fmt.Fprintf(writter, "Hello %s", name)
	}
}

func TestQueryParameter(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Jhon", nil)
	rec := httptest.NewRecorder()

	SayHello(rec, req)

	res := rec.Result()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}

func MultipleQueryParameter(writter http.ResponseWriter, request *http.Request) {
	firstName := request.URL.Query().Get("first_name")
	lastName := request.URL.Query().Get("last_name")

	fmt.Fprintf(writter, "Hello %s %s", firstName, lastName)
}

func TestMultipleQueryParameter(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?first_name=Jhon&last_name=Doe", nil)
	rec := httptest.NewRecorder()

	MultipleQueryParameter(rec, req)

	res := rec.Result()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}

func MultipleParameterValues(writter http.ResponseWriter, request *http.Request) {
	names := request.URL.Query()["name"]

	fmt.Fprint(writter, strings.Join(names, " "))
}

func TestMultipleParameterValues(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Jhon&name=Doe", nil)
	rec := httptest.NewRecorder()

	MultipleParameterValues(rec, req)

	res := rec.Result()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}
