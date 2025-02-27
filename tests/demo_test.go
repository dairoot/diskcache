package cache_test

import (
	"github.com/dairoot/diskcache"

	"testing"
)

func TestDemo(t *testing.T) {
	cacheKey := "cache_key"
	valueStr := "value"
	cache := diskcache.NewDiskCache("../.cache/")

	// Add another item that never expires.
	cache.Set(cacheKey, valueStr, 0)

	// Set the item in the cache if it doesn't already exist.
	cache.SetNx(cacheKey, valueStr, 0)

	// Get the item from the cache.
	cache.Get(cacheKey)

	// Remove the item from the cache.
	cache.Del(cacheKey)

	// Check if the item exists in the cache.
	cache.Exists(cacheKey)

	// Set an item that expires in 60 seconds.
	cache.Expire(cacheKey, 60)

	// add a list item
	cache.LPush(cacheKey, "xx1")
	cache.LPush(cacheKey, "xx2")
	cache.LPush(cacheKey, "xx3")

	// Returns the length of the list stored at key
	cache.LLen(cacheKey)

	// Removes and returns the first element of the list
	cache.LPop(cacheKey)

	// Removes and returns the last element of the list
	cache.RPop(cacheKey)

	// Returns elements within the specified interval in the list
	cache.LRange(cacheKey, 0, 3)

	// Add the specified members to the set stored at key
	cache.SAdd(cacheKey, "xx1")

	// removes and returns one random elements from the set value stored at key
	cache.SPop(cacheKey)

	// Increment the integer value of a key by one
	cache.Incr(cacheKey)

}
