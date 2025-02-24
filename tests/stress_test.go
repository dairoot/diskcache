package cache_test

import (
	"strconv"
	"testing"

	"github.com/dairoot/diskcache"
)

func TestStress(t *testing.T) {
	var valueStr = "value"
	var StressCount = 10000
	cache := diskcache.NewDiskCache("../.cache/")

	for i := 0; i < StressCount; i++ {
		var cacheKey = "TestStress" + strconv.Itoa(i)
		cache.Set(cacheKey, valueStr, 1)
	}

	for i := 0; i < StressCount; i++ {
		var cacheKey = "TestStress" + strconv.Itoa(i)
		cache.Get(cacheKey)
	}

	for i := 0; i < StressCount; i++ {
		var cacheKey = "TestStress" + strconv.Itoa(i)
		cache.Del(cacheKey)
	}

	for i := 0; i < StressCount; i++ {
		var valueStr = "TestStress" + strconv.Itoa(i)
		cache.LPush("TestStress", valueStr)
	}

}
