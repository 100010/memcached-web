package mem

import (
	"github.com/bradfitz/gomemcache/memcache"
)

var Cache *memcache.Client = InitCache("127.0.0.1", "11211")

func InitCache(host string, port string) *memcache.Client {
    endpoint := host + ":" + port
	mc := memcache.New(endpoint)
	return mc
}
