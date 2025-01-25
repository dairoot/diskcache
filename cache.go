package diskcache

import (
	"os"

	"github.com/dairoot/diskcache/api"
)

// NewDiskCache 创建一个新的磁盘缓存实例
func NewDiskCache(baseDir string) (*api.DiskCache, error) {
	// 确保基础目录存在
	if err := os.MkdirAll(baseDir, 0755); err != nil {
		return nil, err
	}
	return &api.DiskCache{
		BaseDir: baseDir,
	}, nil
}
