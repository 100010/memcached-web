package mem

import (
	"github.com/bradfitz/gomemcache/memcache"
)


func InitCache(host string, port string) *memcache.Client {
    endpoint := host + ":" + port
	mc := memcache.New(endpoint)
	return mc
}
