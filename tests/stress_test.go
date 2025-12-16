package cache_test

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/dairoot/diskcache"
)

func TestStress(t *testing.T) {
	var valueStr = "value"
	var StressCount = 10000
	cache := diskcache.NewDiskCache("../.cache/")
	var wg sync.WaitGroup

	// 并发 Set 操作
	start := time.Now()
	for i := range 10000 {
		var cacheKey = "TestStress" + strconv.Itoa(i)
		wg.Add(1)
		go func() {
			defer wg.Done()
			cache.Set(cacheKey, strconv.Itoa(i), 1)
		}()
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("并发 Set 操作 (%d 次): %v\n", 10000, elapsed)

	// 顺序 Set 操作
	start = time.Now()
	for i := range StressCount {
		var cacheKey = "TestStress" + strconv.Itoa(i)
		cache.Set(cacheKey, valueStr, 1)
	}
	elapsed = time.Since(start)
	fmt.Printf("顺序 Set 操作 (%d 次): %v\n", StressCount, elapsed)

	// Get 操作
	start = time.Now()
	for i := range StressCount {
		var cacheKey = "TestStress" + strconv.Itoa(i)
		cache.Get(cacheKey)
	}
	elapsed = time.Since(start)
	fmt.Printf("Get 操作 (%d 次): %v\n", StressCount, elapsed)

	// Del 操作
	start = time.Now()
	for i := range StressCount {
		var cacheKey = "TestStress" + strconv.Itoa(i)
		cache.Del(cacheKey)
	}
	elapsed = time.Since(start)
	fmt.Printf("Del 操作 (%d 次): %v\n", StressCount, elapsed)

	// LPush 操作
	start = time.Now()
	for i := range StressCount {
		var valueStr = "TestStress" + strconv.Itoa(i)
		cache.LPush("TestStress", valueStr)
	}
	elapsed = time.Since(start)
	fmt.Printf("LPush 操作 (%d 次): %v\n", StressCount, elapsed)

}
