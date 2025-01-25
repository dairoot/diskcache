package cache_test

import (
	"encoding/json"
	"testing"

	"github.com/dairoot/diskcache"
)

func TestDel(t *testing.T) {
	var cacheKey = "TestDel"
	cache, _ := diskcache.NewDiskCache("../.cache")

	data := map[string]interface{}{"name": "Tom"}
	value, _ := json.Marshal(data)
	valueStr := string(value)

	cache.Set(cacheKey, valueStr, 1)
	cache.Del(cacheKey)
	_, err := cache.Get(valueStr)
	if err == nil {
		t.Fatal("del failed")
	}

}
