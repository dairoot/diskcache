package diskcache

import (
	"github.com/dairoot/diskcache/api"
)

// NewDiskCache 创建默认 8 分片的 DiskCache
func NewDiskCache(baseDir string) *api.ShardedDiskCache {
	return api.CreateShardedConn(baseDir, 1)
}

// NewDiskCacheWithShards 创建自定义分片数的 DiskCache
func NewDiskCacheWithShards(baseDir string, numShards int) *api.ShardedDiskCache {
	return api.CreateShardedConn(baseDir, numShards)
}
