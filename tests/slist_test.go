package cache_test

import (
	"testing"

	"github.com/dairoot/diskcache"
)

func TestSList(t *testing.T) {
	cache := diskcache.NewDiskCache("../.cache/")
	cacheKey := "TestList"
	cache.Del(cacheKey)
	vaulesList := []string{"1", "2", "3", "4", "5", "1", "2", "3", "4", "5"}
	uniqueMap := make(map[string]struct{})
	for _, value := range vaulesList {
		uniqueMap[value] = struct{}{} // 只存储键，值可以为空结构体
	}

	for _, value := range vaulesList {
		err := cache.SAdd(cacheKey, value)
		if err != nil {
			t.Fatal(err)
		}
	}

	if cache.LLen(cacheKey) != int64(len(uniqueMap)) {
		t.Fatal("SAdd API failed")
	}

	cache.SPop(cacheKey)
	if cache.LLen(cacheKey) != int64(len(uniqueMap)-1) {
		t.Fatal("SPop API failed")
	}

	cache.SRem(cacheKey, "1")

}
