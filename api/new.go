package api

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/exec"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func CreateDiskCacheConn(baseDir string) *DiskCache {
	os.MkdirAll(baseDir, os.ModePerm)
	cmd := exec.Command("sqlite3", "--version")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
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
	_, err := dc.Conn.ExecContext(dc.Ctx, "PRAGMA journal_mode = WAL;")
	if err != nil {
		log.Fatal(err)
	}

	_, err = dc.Conn.ExecContext(dc.Ctx, `
		CREATE TABLE IF NOT EXISTS cache_key (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			key BLOB  NOT NULL,
			access_count INTEGER DEFAULT 0,
			expire_time REAL,
			access_time REAL NOT NULL,
			create_time REAL NOT NULL
		);`)

	if err != nil {
		log.Fatal(err)
	}

	_, err = dc.Conn.ExecContext(dc.Ctx, `
		CREATE TABLE IF NOT EXISTS cache_value (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			key_id INTEGER  NOT NULL,
			value_md5 TEXT,
			value BLOB  NOT NULL
		);`)

	if err != nil {
		log.Fatal(err)
	}

	_, err = dc.Conn.ExecContext(dc.Ctx, `CREATE UNIQUE INDEX IF NOT EXISTS idx_cache_key_key ON cache_key(key);`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = dc.Conn.ExecContext(dc.Ctx, `
		CREATE INDEX IF NOT EXISTS idx_cache_key_expire_time ON cache_key(expire_time) WHERE expire_time IS NOT NULL;
	`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = dc.Conn.ExecContext(dc.Ctx, `
		CREATE INDEX IF NOT EXISTS idx_cache_value_value_md5 ON cache_value(value_md5) WHERE value_md5 IS NOT NULL;
	`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = dc.Conn.ExecContext(dc.Ctx, `
		CREATE INDEX IF NOT EXISTS idx_cache_value_key_id ON cache_value(key_id);
	`)
	if err != nil {
		log.Fatal(err)
	}

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
		time.Sleep(10 * time.Millisecond)
	}

}
