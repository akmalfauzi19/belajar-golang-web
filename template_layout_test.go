package belajar_golang_web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func TemplateLayout(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/header.gohtml", "./templates/layouts.gohtml", "./templates/footer.gohtml"))
	t.ExecuteTemplate(w, "layout", map[string]interface{}{
		"Title": "Template Layout",
		"Name":  "Wick",
	})
}

func TestTamplateLayout(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	rec := httptest.NewRecorder()

	TemplateLayout(rec, req)

	body, _ := ioutil.ReadAll(rec.Result().Body)
	fmt.Println(string(body))

}
