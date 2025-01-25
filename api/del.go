package api

import (
	"os"
	"path/filepath"
)

// Del 删除键值对
func (dc *DiskCache) Del(key string) error {
	dc.mutex.Lock()
	defer dc.mutex.Unlock()

	dirPath, fileName := dc.getKeyPath(key)
	return os.Remove(filepath.Join(dirPath, fileName))
}
