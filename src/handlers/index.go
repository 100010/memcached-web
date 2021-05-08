package handlers

import (
  "net/http"
  "html/template"
  "os"
)

var tpl *template.Template

func Index(w http.ResponseWriter, r *http.Request) {
    wd, _ := os.Getwd()
    tpl = template.Must(template.ParseFiles(wd + "/src/template/index.html"))
}
