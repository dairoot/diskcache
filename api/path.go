package api

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
)

// getKeyPath 获取键的哈希路径
func (dc *DiskCache) getKeyPath(key string) (string, string) {
	hasher := md5.New()
	hasher.Write([]byte(key))
	hashStr := hex.EncodeToString(hasher.Sum(nil))

	dirName := hashStr[:2]
	fileName := hashStr[2:]

	dirPath := filepath.Join(dc.BaseDir, "keys", dirName)
	_ = os.MkdirAll(dirPath, 0755)

	return dirPath, fileName
}

// getValuePath 获取值文件的存储路径
func (dc *DiskCache) getValuePath(key string, value []byte) (string, string, string) {
	data := map[string]string{
		key: string(value),
	}
	jsonData, _ := json.Marshal(data)
	hasher := md5.New()
	hasher.Write(jsonData)
	hashStr := hex.EncodeToString(hasher.Sum(nil))

	dirName := hashStr[:2]
	fileName := hashStr[2:]

	dirPath := filepath.Join(dc.BaseDir, "values", dirName)
	_ = os.MkdirAll(dirPath, 0755)

	return dirPath, fileName, hashStr
}
