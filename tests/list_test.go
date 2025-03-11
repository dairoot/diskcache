package cache_test

import (
	"strings"
	"testing"
	"time"

	"github.com/dairoot/diskcache"
)

func reverse(slice []string) []string {
	reversed := make([]string, len(slice))

	for i, v := range slice {
		reversed[len(slice)-1-i] = v
	}
	return reversed
}
func TestList(t *testing.T) {
	cache := diskcache.NewDiskCache("../.cache/")
	cacheKey := "TestList"
	cache.Del(cacheKey)
	vaulesList := []string{"1", "2", "3", "4", "5"}
	for _, value := range vaulesList {
		err := cache.LPush(cacheKey, value)
		if err != nil {
			t.Fatal(err)
		}
	}

	if cache.LLen(cacheKey) != int64(len(vaulesList)) {
		t.Fatal("LLen API failed")
	}
	llen := cache.LLen(cacheKey)

	cache_list := cache.LRange(cacheKey, 0, llen)
	if strings.Join(cache_list, ",") != strings.Join(reverse(vaulesList)[0:llen], ",") {
		t.Fatal("LRange API failed")
	}

	cache_list = cache.RRange(cacheKey, 0, llen)
	if strings.Join(cache_list, ",") != strings.Join(vaulesList[0:llen], ",") {
		t.Fatal("RRange API failed")
	}

	cacheData, err := cache.LPop(cacheKey)
	if cacheData != vaulesList[len(vaulesList)-1] {
		t.Fatal("LPop API failed: ", err)
	}

	cacheData, err = cache.RPop(cacheKey)
	if cacheData != vaulesList[0] {
		t.Fatal("RPop API failed: ", err)
	}

	cache_listv1 := cache.LRange(cacheKey, 0, llen)
	cache.LRem(cacheKey, cache_listv1[0])
	cache_listv2 := cache.LRange(cacheKey, 0, llen)
	if strings.Join(cache_listv1[1:], ",") != strings.Join(cache_listv2, ",") {
		t.Fatal("LRange API failed")
	}

	cache.Expire(cacheKey, 0.1)
	time.Sleep(time.Millisecond * (0.2 * 1000))
	if cache.Exists(cacheKey) == true {
		t.Fatal("expire failed")
	}

}
