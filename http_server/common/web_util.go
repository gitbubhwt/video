package common

import (
	"html/template"
	"net/http"
)

//跳转页面
func GoToPage(w http.ResponseWriter, htmlPath string,data interface{}) {
	if t, err := template.ParseFiles(htmlPath); err == nil {
		t.Execute(w, data)
	}
}
