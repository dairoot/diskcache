package diskcache

import (
	"time"

	"github.com/dairoot/diskcache/api"
)

func NewDiskCache(baseDir string) *api.DiskCache {
	cache := api.CreateDiskCacheConn(baseDir)
	cache.InitDb()
	_ = cache.DelExpire()
	cache.StartMaintenance(5 * time.Minute)
	return cache
}
