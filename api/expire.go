package api

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

// Expire 设置键的过期时间并返回当前值
func (dc *DiskCache) Expire(key string, ttl int64) error {
	dc.mutex.Lock()
	defer dc.mutex.Unlock()

	// 获取键的元数据
	item, err := dc.getKeyInfo(key)
	if err != nil {
		return err
	}

	// 更新过期时间
	item.TTL = ttl
	item.Time = time.Now().Unix()

	// 保存更新后的元数据
	newData, err := json.Marshal(item)
	if err != nil {
		return err
	}

	dirPath, fileName := dc.getKeyPath(key)
	return os.WriteFile(filepath.Join(dirPath, fileName), newData, 0644)

}
