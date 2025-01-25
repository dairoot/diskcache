package api

import (
	"os"
	"path/filepath"
)

// Exists 检查键是否存在
func (dc *DiskCache) Exists(key string) bool {
	dc.mutex.RLock()
	defer dc.mutex.RUnlock()

	dirPath, fileName := dc.getKeyPath(key)
	_, err := os.Stat(filepath.Join(dirPath, fileName))
	return err == nil
}
