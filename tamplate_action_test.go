package belajar_golang_web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func TemplateActionIf(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("./templates/if.gohtml"))
	t.ExecuteTemplate(w, "if.gohtml", Page{
		Title: "Template Data Struct",
		Name:  "wick",
	})
}

func TestTemplateActionIf(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	rec := httptest.NewRecorder()

	TemplateActionIf(rec, req)

	body, _ := ioutil.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionOperator(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("./templates/comparator.gohtml"))
	t.ExecuteTemplate(w, "comparator.gohtml", map[string]interface{}{
		"Title":      "Template Data Map",
		"FinalValue": 100,
	})
}

func TestTemplateActionOperator(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	rec := httptest.NewRecorder()

	TemplateActionOperator(rec, req)

	body, _ := ioutil.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionRange(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("./templates/range.gohtml"))
	t.ExecuteTemplate(w, "range.gohtml", map[string]interface{}{
		"Title": "Template Data Map",
		"Items": []string{"wick", "jane", "jason"},
	})
}

func TestTemplateActionRange(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	rec := httptest.NewRecorder()

	TemplateActionRange(rec, req)

	body, _ := ioutil.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionWith(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("./templates/address.gohtml"))
	t.ExecuteTemplate(w, "address.gohtml", map[string]interface{}{
		"Title": "Template Data Map",
		"Name":  "Wick",
		"Address": map[string]interface{}{
			"Street": "Jl. Sudirman",
			"City":   "Jakarta",
			"State":  "DKI Jakarta",
			"Zip":    "12",
		},
	})
}

func TestTemplateActionWith(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	rec := httptest.NewRecorder()

	TemplateActionWith(rec, req)

	body, _ := ioutil.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}
