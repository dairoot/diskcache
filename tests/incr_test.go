package cache_test

import (
	"testing"

	"github.com/dairoot/diskcache"
)

func TestIncr(t *testing.T) {
	var cacheKey = "TestIncr"
	cache := diskcache.NewDiskCache("../.cache/")
	cache.Del(cacheKey)

	_ = cache.Incr(cacheKey)

	cache_value := cache.Incr(cacheKey)
	if cache_value != 1 {
		t.Fatal("Incr API failed")
	}

}
