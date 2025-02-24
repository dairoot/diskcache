package diskcache

import (
	"github.com/dairoot/diskcache/api"
)

func NewDiskCache(baseDir string) *api.DiskCache {
	cache := api.CreateDiskCacheConn(baseDir)
	cache.InitDb()
	return cache
}
