package cache_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/dairoot/diskcache"
)

func TestGet(t *testing.T) {
	var cacheKey = "TestGet"
	cache, _ := diskcache.NewDiskCache("../.cache")

	data := map[string]interface{}{"name": "Tom"}
	value, _ := json.Marshal(data)
	valueStr := string(value)

	cache.Set(cacheKey, valueStr, 2)

	// 获取值
	_, err := cache.Get(cacheKey)
	if err != nil {
		t.Fatal(err)
	}

	// 重新获取值，验证是否过期
	time.Sleep(time.Second * 3)
	_, err = cache.Get(cacheKey)
	if err == nil {
		t.Fatal("set cache time failed")
	}

}
