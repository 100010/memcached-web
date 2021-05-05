package client

import (
  "fmt"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, World")
}

var port string
func ServerInit() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8081", nil)
}
