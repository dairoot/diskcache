# DiskCache: Disk Backed Cache

### Installation

```sh
go get github.com/dairoot/diskcache
```

### Usage

```go
package main

import (
	"github.com/dairoot/diskcache"
)

func main() {
	cacheKey := "cache_key"
	valueStr := "value"

	cache, _ := diskcache.NewDiskCache("./.cache")

	// Add another item that never expires.
	cache.Set(cacheKey, valueStr, 0)

	// Get the item from the cache.
	cache.Get(cacheKey)

	// Remove the item from the cache.
	cache.Del(cacheKey)

	// Check if the item exists in the cache.
	cache.Exists(cacheKey)

	// Set an item that expires in 60 seconds.
	cache.Expire(cacheKey, 60)
}
```
