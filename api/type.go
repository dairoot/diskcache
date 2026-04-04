package api

import (
	"context"
	"database/sql"
)

type DiskCache struct {
	Ctx    context.Context
	DB     *sql.DB
	Conn   *sql.Conn
	stopCh chan struct{}
}

// Close 停止后台维护并关闭数据库连接
func (dc *DiskCache) Close() error {
	if dc.stopCh != nil {
		close(dc.stopCh)
	}
	if dc.Conn != nil {
		dc.Conn.Close()
	}
	if dc.DB != nil {
		return dc.DB.Close()
	}
	return nil
}
