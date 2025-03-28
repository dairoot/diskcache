package cache_test

import (
	"testing"

	"github.com/dairoot/diskcache"
)

func TestSet(t *testing.T) {
	cache := diskcache.NewDiskCache("../.cache/")
	var valueStr = "value"

	err := cache.Set("TestSet1", valueStr, 1)
	if err != nil {
		t.Fatal(err)
	}

	err = cache.Set("TestSet2", valueStr, 1)
	if err != nil {
		t.Fatal(err)
	}

}
