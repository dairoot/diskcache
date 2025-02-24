package cache_test

import (
	"testing"

	"github.com/dairoot/diskcache"
)

func TestDel(t *testing.T) {
	var cacheKey = "TestDel"
	var valueStr = "value"

	cache := diskcache.NewDiskCache("../.cache/")

	cache.Set(cacheKey, valueStr, 1)
	cache.Del(cacheKey)
	_, err := cache.Get(valueStr)
	if err == nil {
		t.Fatal("del failed")
	}

}
