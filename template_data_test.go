package belajar_golang_web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func TemplateDataMap(w http.ResponseWriter, r *http.Request) {
	// t := template.Must(template.New("data").Parse(`{{.Name}} {{.Age}}`))
	// t.ExecuteTemplate(w, "data", map[string]interface{}{
	// 	"Name": "wick",
	// 	"Age":  27,
	// })
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(w, "name.gohtml", map[string]interface{}{
		"Title": "Template Data Map",
		"Name":  "wick",
		"Address": map[string]interface{}{
			"City": "KOTA Bandung",
		},
	})
}

func TestTemplateDataMap(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	rec := httptest.NewRecorder()

	TemplateDataMap(rec, req)

	body, _ := ioutil.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}

type Address struct {
	City string
}

type Page struct {
	Title   string
	Name    string
	Address Address
}

func TemplateDataStruct(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(w, "name.gohtml", Page{
		Title: "Template Data Struct",
		Name:  "wick",
		Address: Address{
			City: "KOTA Bandung",
		},
	})
}

func TestTemplateDataStruct(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	rec := httptest.NewRecorder()

	TemplateDataStruct(rec, req)

	body, _ := ioutil.ReadAll(rec.Result().Body)
	fmt.Println(string(body))
}
