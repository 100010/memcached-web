package handlers

import (
  "net/http"
  "html/template"
  "os"
)

var tpl *template.Template

func Index(w http.ResponseWriter, r *http.Request) {
    wd, _ := os.Getwd()
    t, err := template.ParseFiles(wd + "/src/template/index.html")
    if err != nil {
        panic(err.Error())
    }
    if err := t.Execute(w, nil); err != nil {
		panic(err.Error())
	}
}
