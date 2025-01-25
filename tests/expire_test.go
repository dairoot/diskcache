package cache_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/dairoot/diskcache"
)

func TestExpire(t *testing.T) {
	var cacheKey = "TestExpire"

	cache, _ := diskcache.NewDiskCache("../.cache")

	data := map[string]interface{}{"name": "Tom"}
	value, _ := json.Marshal(data)
	valueStr := string(value)

	cache.Set(cacheKey, valueStr, 0)

	cache.Expire(cacheKey, 1)

	time.Sleep(time.Second * 2)

	_, err := cache.Get(cacheKey)
	if err == nil {
		t.Fatal("expire failed")
	}

}
