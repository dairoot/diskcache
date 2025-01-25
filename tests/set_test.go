package cache_test

import (
	"encoding/json"
	"testing"

	"github.com/dairoot/diskcache"
)

func TestSet(t *testing.T) {
	cache, _ := diskcache.NewDiskCache("../.cache")

	data := map[string]interface{}{"name": "Tom"}
	value, _ := json.Marshal(data)
	valueStr := string(value)

	err := cache.Set("TestSet1", valueStr, 1)
	if err != nil {
		t.Fatal(err)
	}

	err = cache.Set("TestSet2", valueStr, 1)
	if err != nil {
		t.Fatal(err)
	}

}
