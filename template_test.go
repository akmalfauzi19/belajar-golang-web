package belajar_golang_web

import (
	"embed"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func SimpleHTML(w http.ResponseWriter, r *http.Request) {
	templateText := `<html><body>{{.}}</body></html>`
	// t, err := template.Must(template.New("SIMPLE").Parse(templateText))
	// if err != nil {
	// 	panic(err)
	// }

	t := template.Must(template.New("SIMPLE").Parse(templateText))

	t.ExecuteTemplate(w, "SIMPLE", "Hello Template")
}

func TestSimpleHTML(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8080", nil)
	rec := httptest.NewRecorder()

	SimpleHTML(rec, req)

	body, _ := ioutil.ReadAll(rec.Result().Body)
	fmt.Println(string(body))

}

func SimpleHTMLFile(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/simple.gohtml"))
	t.ExecuteTemplate(w, "simple.gohtml", "Hello Template ggwp")
}

func TestSimpleHTMLFile(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8080", nil)
	rec := httptest.NewRecorder()

	SimpleHTMLFile(rec, req)

	body, _ := ioutil.ReadAll(rec.Result().Body)
	fmt.Println(string(body))

}

func TemplateDirectory(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))
	t.ExecuteTemplate(w, "simple.gohtml", "Hello Template ggwp")
}

func TestTemplateDirectory(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8080", nil)
	rec := httptest.NewRecorder()

	TemplateDirectory(rec, req)

	body, _ := ioutil.ReadAll(rec.Result().Body)
	fmt.Println(string(body))

}

//go:embed templates/*.gohtml
var templates embed.FS

func TemplateEmbed(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))
	t.ExecuteTemplate(w, "simple.gohtml", "Hello Template ggwp")
}

func TestTemplateEmbed(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8080", nil)
	rec := httptest.NewRecorder()

	TemplateEmbed(rec, req)

	body, _ := ioutil.ReadAll(rec.Result().Body)
	fmt.Println(string(body))

}
