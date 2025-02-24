package cache_test

import (
	"testing"
	"time"

	"github.com/dairoot/diskcache"
)

func TestExpire(t *testing.T) {
	var cacheKey = "TestExpire"
	var valueStr = "value"

	cache := diskcache.NewDiskCache("../.cache/")

	cache.Set(cacheKey, valueStr, 0)

	cache.Expire(cacheKey, 0.1)

	time.Sleep(time.Millisecond * (0.2 * 1000))

	if cache.Exists(cacheKey) == true {
		t.Fatal("expire failed")
	}

}
