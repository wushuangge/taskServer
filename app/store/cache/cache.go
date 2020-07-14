package cache

import "C"
import (
	"github.com/patrickmn/go-cache"
	"time"
)

//内存缓存
var c *cache.Cache

func InitMemoryCache() {
	c = cache.New(time.Second, time.Minute)
}
