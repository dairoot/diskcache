package cache_test

import (
	"testing"
	"time"

	"github.com/dairoot/diskcache"
)

func TestSetNx(t *testing.T) {
	cache := diskcache.NewDiskCache("../.cache/")
	var valueStr = "value"

	cache.Del("TestSetNx")
	isInsert, _ := cache.SetNx("TestSetNx", valueStr, 0.1)
	if isInsert != 1 {
		t.Fatal("SetNx API failed")
	}
	time.Sleep(0.3 * 1000 * time.Millisecond)
	isInsert, _ = cache.SetNx("TestSetNx", valueStr, 5)
	if isInsert != 1 {
		t.Fatal("SetNx API failed")
	}

	isInsert, _ = cache.SetNx("TestSetNx", valueStr, 5)
	if isInsert == 1 {
		t.Fatal("SetNx API failed")
	}
}
