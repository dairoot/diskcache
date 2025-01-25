package api

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// getKeyInfo 获取键的元数据信息
func (dc *DiskCache) getKeyInfo(key string) (*CacheItem, error) {
	dirPath, fileName := dc.getKeyPath(key)
	data, err := os.ReadFile(filepath.Join(dirPath, fileName))
	if err != nil {
		return nil, fmt.Errorf("key not found")
	}

	var item CacheItem
	if err := json.Unmarshal(data, &item); err != nil {
		return nil, err
	}

	// 检查是否过期
	if item.TTL > 0 {
		if time.Now().Unix() > item.Time+item.TTL {
			go dc.Del(key)
			return nil, fmt.Errorf("key expired")
		}
	}

	return &item, nil
}

// Get 获取键对应的值
func (dc *DiskCache) Get(key string) (string, error) {
	dc.mutex.RLock()
	defer dc.mutex.RUnlock()

	// 获取键的元数据
	item, err := dc.getKeyInfo(key)
	if err != nil {
		return "", err
	}

	// 读取值文件
	valueDirPath := filepath.Join(dc.BaseDir, "values", item.ValueHash[:2])
	valueData, err := os.ReadFile(filepath.Join(valueDirPath, item.ValueHash[2:]))
	if err != nil {
		return "", fmt.Errorf("value file not found")
	}

	return string(valueData), nil
}
