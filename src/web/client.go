package client

import (
  "net/http"
  "github.com/100010/memcached-web/src/handlers"
  "github.com/100010/memcached-web/src/mem"
)

func ServerInit() {
    http.HandleFunc("/", handlers.Index)
    http.ListenAndServe(":8081", nil)
    mem.Cache.GetMulti()
}
