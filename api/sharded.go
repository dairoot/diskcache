package api

// Get 从对应分片获取值
func (s *ShardedDiskCache) Get(cacheKey string) (string, error) {
	return s.shard(cacheKey).Get(cacheKey)
}

// Exists 检查 key 是否存在
func (s *ShardedDiskCache) Exists(cacheKey string) bool {
	return s.shard(cacheKey).Exists(cacheKey)
}

// Set 设置 key-value，ttl 为过期秒数（0 表示永不过期）
func (s *ShardedDiskCache) Set(cacheKey string, cacheValue string, ttl float64) error {
	return s.shard(cacheKey).Set(cacheKey, cacheValue, ttl)
}

// SetNx 若 key 不存在则插入，返回是否为新插入
func (s *ShardedDiskCache) SetNx(cacheKey string, cacheValue string, ttl float64) (int16, error) {
	return s.shard(cacheKey).SetNx(cacheKey, cacheValue, ttl)
}

// Del 删除指定 key
func (s *ShardedDiskCache) Del(cacheKey string) error {
	return s.shard(cacheKey).Del(cacheKey)
}

// DelExpire 清理所有分片中的过期 key
func (s *ShardedDiskCache) DelExpire() error {
	var lastErr error
	for _, dc := range s.shards {
		if err := dc.DelExpire(); err != nil {
			lastErr = err
		}
	}
	return lastErr
}

// Expire 设置 key 的过期时间
func (s *ShardedDiskCache) Expire(cacheKey string, ttl float64) error {
	return s.shard(cacheKey).Expire(cacheKey, ttl)
}

// Incr 对 key 执行原子自增
func (s *ShardedDiskCache) Incr(cacheKey string) int64 {
	return s.shard(cacheKey).Incr(cacheKey)
}

// LPush 向列表头部追加元素
func (s *ShardedDiskCache) LPush(cacheKey string, cacheValue string) error {
	return s.shard(cacheKey).LPush(cacheKey, cacheValue)
}

// LPop 从列表头部弹出元素
func (s *ShardedDiskCache) LPop(cacheKey string) (string, error) {
	return s.shard(cacheKey).LPop(cacheKey)
}

// RPop 从列表尾部弹出元素
func (s *ShardedDiskCache) RPop(cacheKey string) (string, error) {
	return s.shard(cacheKey).RPop(cacheKey)
}

// LRange 从列表头部取范围元素
func (s *ShardedDiskCache) LRange(cacheKey string, offset int64, limit int64) []string {
	return s.shard(cacheKey).LRange(cacheKey, offset, limit)
}

// RRange 从列表尾部取范围元素
func (s *ShardedDiskCache) RRange(cacheKey string, offset int64, limit int64) []string {
	return s.shard(cacheKey).RRange(cacheKey, offset, limit)
}

// LLen 获取列表长度
func (s *ShardedDiskCache) LLen(cacheKey string) int64 {
	return s.shard(cacheKey).LLen(cacheKey)
}

// LRem 从列表中删除指定值
func (s *ShardedDiskCache) LRem(cacheKey string, cacheValue string) error {
	return s.shard(cacheKey).LRem(cacheKey, cacheValue)
}

// SAdd 向集合中添加元素
func (s *ShardedDiskCache) SAdd(cacheKey string, cacheValue string) error {
	return s.shard(cacheKey).SAdd(cacheKey, cacheValue)
}

// SPop 从集合中随机弹出一个元素
func (s *ShardedDiskCache) SPop(cacheKey string) (string, error) {
	return s.shard(cacheKey).SPop(cacheKey)
}

// SRem 从集合中删除指定元素
func (s *ShardedDiskCache) SRem(cacheKey string, cacheValue string) error {
	return s.shard(cacheKey).SRem(cacheKey, cacheValue)
}

// Vacuum 对所有分片执行增量回收
func (s *ShardedDiskCache) Vacuum() {
	for _, dc := range s.shards {
		dc.Vacuum()
	}
}
