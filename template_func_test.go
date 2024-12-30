package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"text/template"
)

type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + ", my name is " + myPage.Name
}

func TemplateFunction(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{ .SayHello "wick" }}`))
	t.ExecuteTemplate(w, "FUNCTION", MyPage{Name: "john"})
}

func TestTemplateFunction(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	res := httptest.NewRecorder()

	TemplateFunction(res, req)

	body, _ := io.ReadAll(res.Result().Body)
	fmt.Println(string(body))
}

func TemplateFunctionGlobal(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{ len .Name }}`))
	t.ExecuteTemplate(w, "FUNCTION", MyPage{Name: "john"})
}

func TestTemplateFunctionGlobal(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	res := httptest.NewRecorder()

	TemplateFunctionGlobal(res, req)

	body, _ := io.ReadAll(res.Result().Body)
	fmt.Println(string(body))
}

func TemplateFunctionCreateGlobal(w http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION")
	t.Funcs(map[string]interface{}{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})

	t = template.Must(t.Parse(`{{ upper .Name }}`))

	t.ExecuteTemplate(w, "FUNCTION", MyPage{Name: "john"})

}

func TestTemplateFunctionCreateGlobal(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	res := httptest.NewRecorder()

	TemplateFunctionCreateGlobal(res, req)

	body, _ := io.ReadAll(res.Result().Body)
	fmt.Println(string(body))
}

func TemplateFunctionCreateGlobalPipeline(w http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION")
	t.Funcs(map[string]interface{}{
		"sayHello": func(name string) string {
			return "Hello " + name
		},
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})

	t = template.Must(t.Parse(`{{ sayHello .Name | upper}}`))

	t.ExecuteTemplate(w, "FUNCTION", MyPage{Name: "john"})
}
func TestTemplateFunctionCreateGlobalPipeline(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	res := httptest.NewRecorder()

	TemplateFunctionCreateGlobalPipeline(res, req)

	body, _ := io.ReadAll(res.Result().Body)
	fmt.Println(string(body))
}
