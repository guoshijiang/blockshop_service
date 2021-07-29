package common

import (
	cache "github.com/patrickmn/go-cache"
	"time"
)

var Che *cache.Cache = cache.New(60*time.Minute, 120*time.Minute)
