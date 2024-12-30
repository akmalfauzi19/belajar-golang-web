package belajar_golang_web

import (
	"embed"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

//go:embed templates/*.gohtml
var templates1 embed.FS

var myTemplates = template.Must(template.ParseFS(templates1, "templates/*.gohtml"))

// variable must place outside the function for caching the template

func templateCaching(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "simple.gohtml", "Hello Template Caching")
}

func TestTemplateCaching(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	res := httptest.NewRecorder()

	templateCaching(res, req)

	body, _ := io.ReadAll(res.Result().Body)
	fmt.Println(string(body))
}
