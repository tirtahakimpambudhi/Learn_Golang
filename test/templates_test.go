package test

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)



var TemplatesCaching = func ()*template.Template{
	tmpl := template.New("index")
	tmpl = tmpl.Funcs(map[string]interface{}{
		"uppercase" : func (value string) interface{} {
			return strings.ToUpper(value)
		},
		"sayhello":func (value string) interface{} {
			return fmt.Sprintf("Hello %v",value)
		},
	})
	tmpl = template.Must(template.ParseGlob("./templates_test/*.gohtml"))
	return tmpl
}()

type TestingData struct {
	Name string
	Title string
}

func GlobFunc(w http.ResponseWriter, r *http.Request) {
	tmpl := template.New("Global_Function")
	tmpl = tmpl.Funcs(map[string]interface{}{
		"upper" : func (value string) interface{} {
			return strings.ToUpper(value)
		},
	})

	tmpl.Parse("{{upper .Name}}")
	tmpl.ExecuteTemplate(w,"Global_Function",TestingData{
		Name: "tirta",
		Title: "testing - global function",
	})
}

func TemplatesCachingCommon(w http.ResponseWriter, r *http.Request) {
	TemplatesCaching.ExecuteTemplate(w,"index.gohtml",TestingData{
		Title: "Testing - Caching",
		Name: "Hello Caching",
	})
}

func TestGlobFunction(t *testing.T) {
	request := httptest.NewRequest("GET","http://localhost:8999/",nil)
	recorder := httptest.NewRecorder()

	GlobFunc(recorder,request)

	response := recorder.Result()
	body , _ := io.ReadAll(response.Body)

	t.Log(string(body))
}
func TestCaching(t *testing.T) {
	request := httptest.NewRequest("GET","http://localhost:8999/",nil)
	recorder := httptest.NewRecorder()

	TemplatesCachingCommon(recorder,request)

	response := recorder.Result()
	body , _ := io.ReadAll(response.Body)

	t.Log(string(body))
}










































