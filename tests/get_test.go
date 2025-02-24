package cache_test

import (
	"testing"

	"github.com/dairoot/diskcache"
)

func TestGet(t *testing.T) {
	var cacheKey = "TestGet"
	var valueStr = "value"

	cache := diskcache.NewDiskCache("../.cache/")

	cache.Set(cacheKey, valueStr, 0)

	// 获取值
	_, err := cache.Get(cacheKey)
	if err != nil {
		t.Fatal(err)
	}

}
