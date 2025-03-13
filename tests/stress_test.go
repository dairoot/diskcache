package cache_test

import (
	"strconv"
	"sync"
	"testing"

	"github.com/dairoot/diskcache"
)

func TestStress(t *testing.T) {
	var valueStr = "value"
	var StressCount = 10000
	cache := diskcache.NewDiskCache("../.cache/")
	var wg sync.WaitGroup

	for i := range 10000 {
		var cacheKey = "TestStress" + strconv.Itoa(i)
		wg.Add(1)
		go func() {
			defer wg.Done()
			cache.Set(cacheKey, strconv.Itoa(i), 1)
		}()
	}
	wg.Wait()

	for i := range StressCount {
		var cacheKey = "TestStress" + strconv.Itoa(i)
		cache.Set(cacheKey, valueStr, 1)
	}

	for i := range StressCount {
		var cacheKey = "TestStress" + strconv.Itoa(i)
		cache.Get(cacheKey)
	}

	for i := range StressCount {
		var cacheKey = "TestStress" + strconv.Itoa(i)
		cache.Del(cacheKey)
	}

	for i := range StressCount {
		var valueStr = "TestStress" + strconv.Itoa(i)
		cache.LPush("TestStress", valueStr)
	}

}
