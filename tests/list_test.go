package cache_test

import (
	"testing"
	"time"

	"github.com/dairoot/diskcache"
)

func TestList(t *testing.T) {
	cache, _ := diskcache.NewDiskCache("../.cache")
	cacheKey := "TestList"
	cache.LPush(cacheKey, "1")
	cache.LPush(cacheKey, "2")
	cache.LPush(cacheKey, "3")

	cache.Expire(cacheKey, 1)
	cache_data, _ := cache.LPop(cacheKey)
	if cache_data != "3" {
		t.Fatal("LPop API failed")
	}

	cache_data, _ = cache.RPop(cacheKey)
	if cache_data != "1" {
		t.Fatal("RPop API failed")
	}

	time.Sleep(time.Second * 2)
	cache_data, _ = cache.LPop(cacheKey)
	if cache_data != "" {
		t.Fatal("expire failed")
	}

}
