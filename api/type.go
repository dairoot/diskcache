package api

import (
	"context"
	"database/sql"
	"hash/fnv"
)

type DiskCache struct {
	Ctx    context.Context
	DB     *sql.DB
	Conn   *sql.Conn
	stopCh chan struct{}
}

// Close 停止后台维护并关闭数据库连接
func (dc *DiskCache) Close() error {
	if dc.stopCh != nil {
		close(dc.stopCh)
	}
	if dc.Conn != nil {
		dc.Conn.Close()
	}
	if dc.DB != nil {
		return dc.DB.Close()
	}
	return nil
}

// ShardedDiskCache 将 key 按 fnv32 哈希路由到多个独立 SQLite 分片
type ShardedDiskCache struct {
	shards    []*DiskCache
	numShards uint32
}

func (s *ShardedDiskCache) shard(key string) *DiskCache {
	h := fnv.New32a()
	h.Write([]byte(key))
	return s.shards[h.Sum32()%s.numShards]
}

// Close 关闭所有分片
func (s *ShardedDiskCache) Close() error {
	var lastErr error
	for _, dc := range s.shards {
		if err := dc.Close(); err != nil {
			lastErr = err
		}
	}
	return lastErr
}
