package api

import (
	"context"
	"database/sql"
)

type DiskCache struct {
	Ctx  context.Context
	DB   *sql.DB
	Conn *sql.Conn
}

// Close 关闭数据库连接，释放资源
func (dc *DiskCache) Close() error {
	if dc.Conn != nil {
		dc.Conn.Close()
	}
	if dc.DB != nil {
		return dc.DB.Close()
	}
	return nil
}
