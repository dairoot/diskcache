package cache_test

import (
	"testing"

	"github.com/dairoot/diskcache"
)

func TestIncr(t *testing.T) {
	var cacheKey = "TestIncr"
	cache := diskcache.NewDiskCache("../.cache/")
	cache.Del(cacheKey)

	// 第一次 Incr，值从 0 变为 1，返回 1
	firstValue := cache.Incr(cacheKey)
	if firstValue != 1 {
		t.Fatalf("First Incr should return 1, got %d", firstValue)
	}

	// 第二次 Incr，值从 1 变为 2，返回 2
	secondValue := cache.Incr(cacheKey)
	if secondValue != 2 {
		t.Fatalf("Second Incr should return 2, got %d", secondValue)
	}
}
