package cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

func Init() *cache.Cache {
	return cache.New(cache.NoExpiration, 1*time.Minute)
}
