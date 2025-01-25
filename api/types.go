package api

import "sync"

// DiskCache 表示磁盘缓存结构
type DiskCache struct {
	BaseDir string
	mutex   sync.RWMutex
}

// CacheItem 表示缓存项的数据结构
type CacheItem struct {
	Key       string `json:"key"`
	Time      int64  `json:"time"`
	TTL       int64  `json:"ttl"` // 过期时间（秒），0表示永不过期
	ValueHash string `json:"value_hash"`
}
