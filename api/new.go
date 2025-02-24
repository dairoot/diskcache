package api

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func CreateDiskCacheConn(baseDir string) *DiskCache {
	os.MkdirAll(baseDir, os.ModePerm)
	db, err := sql.Open("sqlite3", baseDir+"/cache.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	conn, _ := db.Conn(ctx)

	return &DiskCache{
		Ctx:  ctx,
		Conn: conn,
	}
}

func (dc *DiskCache) InitDb() {
	// 初始化表
	dc.Conn.ExecContext(dc.Ctx, "PRAGMA journal_mode = WAL;")

	dc.Conn.ExecContext(dc.Ctx, `
		CREATE TABLE IF NOT EXISTS cache_key (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			key BLOB  NOT NULL,
			access_count INTEGER DEFAULT 0,
			expire_time REAL,
			access_time REAL NOT NULL,
			create_time REAL NOT NULL
		);`)

	dc.Conn.ExecContext(dc.Ctx, `
		CREATE TABLE IF NOT EXISTS cache_value (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			key_id INTEGER  NOT NULL,
			value BLOB  NOT NULL
		);`)

	dc.Conn.ExecContext(dc.Ctx, `
		CREATE INDEX IF NOT EXISTS idx_cache_key_key ON cache_key(key);
	`)

	dc.Conn.ExecContext(dc.Ctx, `
		CREATE INDEX IF NOT EXISTS idx_cache_key_expire_time ON cache_key(expire_time) WHERE expire_time IS NOT NULL
	`)

	dc.Conn.ExecContext(dc.Ctx, `
		CREATE INDEX IF NOT EXISTS idx_cache_value_key_id ON cache_value(key_id);
	`)

}
func (dc *DiskCache) Tx() *sql.Tx {

	// tx, err := dc.Conn.BeginTx(dc.Ctx, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// return tx

	for {
		tx, err := dc.Conn.BeginTx(dc.Ctx, nil)
		if err == nil {
			return tx
		}
		time.Sleep(1 * time.Millisecond)
	}

}
