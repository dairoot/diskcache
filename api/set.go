package api

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

// Set 设置键值对
func (dc *DiskCache) Set(key string, value string, ttl int64) error {

	dc.mutex.Lock()
	defer dc.mutex.Unlock()
	valueByte := []byte(value)
	valueDirPath, valueFileName, valueHash := dc.getValuePath(key, valueByte)
	valueFilePath := filepath.Join(valueDirPath, valueFileName)

	if _, err := os.Stat(valueFilePath); os.IsNotExist(err) {
		if err := os.WriteFile(valueFilePath, valueByte, 0644); err != nil {
			return err
		}
	}

	item := CacheItem{
		Key:       key,
		Time:      time.Now().Unix(),
		TTL:       ttl,
		ValueHash: valueHash,
	}

	data, err := json.Marshal(item)
	if err != nil {
		return err
	}

	dirPath, fileName := dc.getKeyPath(key)
	return os.WriteFile(filepath.Join(dirPath, fileName), data, 0644)
}
