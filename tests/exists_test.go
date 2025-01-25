package cache_test

import (
	"encoding/json"
	"testing"

	"github.com/dairoot/diskcache"
)

func TestExists(t *testing.T) {
	var cacheKey = "TestExists"

	cache, _ := diskcache.NewDiskCache("../.cache")

	data := map[string]interface{}{"name": "Tom"}
	value, _ := json.Marshal(data)
	valueStr := string(value)

	cache.Set(cacheKey, valueStr, 0)

	// 检查是否存在
	is_exists := cache.Exists(cacheKey)
	if !is_exists {
		t.Fatal("Exists API failed")
	}

	is_exists = cache.Exists("TestExpire2")
	if is_exists {
		t.Fatal("Exists API failed")
	}

}
