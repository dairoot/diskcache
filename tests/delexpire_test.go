package cache_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/dairoot/diskcache"
)

func TestDelExpire(t *testing.T) {
	cacheKey := "TestDelExpire"
	valueStr := "value"

	cache := diskcache.NewDiskCache("../.cache/")
	for i := 0; i < 10; i++ {
		cache.Set(cacheKey+strconv.Itoa(i), valueStr, 0.01)
	}
	time.Sleep(time.Millisecond * (0.2 * 1000))
	cache.DelExpire()

}
