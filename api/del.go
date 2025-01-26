package api

import (
	"os"
	"path/filepath"
)

// delNotLock 删除 key 文件
func (dc *DiskCache) delKeyFile(key string) {
	dirPath, fileName := dc.getKeyPath(key)
	os.Remove(filepath.Join(dirPath, fileName))

}

// delValueFile 删除value文件
func (dc *DiskCache) delValueFile(valueHash string) {
	oldValueDirPath := filepath.Join(dc.BaseDir, "values", valueHash[:2])
	oldValuePath := filepath.Join(oldValueDirPath, valueHash[2:])
	_ = os.Remove(oldValuePath)
}

// Del 删除键值对
func (dc *DiskCache) Del(key string) error {
	dc.mutex.Lock()
	defer dc.mutex.Unlock()
	item, err := dc.getKeyInfo(key)
	if err != nil {
		return err
	}

	dc.delKeyFile(key)
	dc.delValueFile(item.ValueHash)
	return nil
}
