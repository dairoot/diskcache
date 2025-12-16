package api

import (
	"time"
)

func (dc *DiskCache) DelExpire() error {
	nowTime := time.Now().Unix()
	rows, err := dc.Conn.QueryContext(dc.Ctx, "SELECT id FROM cache_key WHERE expire_time <= ? order by access_time asc limit 100", nowTime)
	if err != nil {
		return err
	}
	defer rows.Close()

	// 遍历结果
	var keyIDs []int64
	for rows.Next() {
		var keyID int64
		if err := rows.Scan(&keyID); err != nil {
			return err
		}
		keyIDs = append(keyIDs, keyID)
	}

	if len(keyIDs) == 0 {
		return nil
	}

	tx := dc.Tx()
	if tx == nil {
		return nil
	}

	// 使用参数化查询删除
	for _, keyID := range keyIDs {
		tx.Exec("DELETE FROM cache_value WHERE key_id = ?", keyID)
		tx.Exec("DELETE FROM cache_key WHERE id = ?", keyID)
	}

	return tx.Commit()
}

func (dc *DiskCache) Del(cacheKey string) error {
	tx := dc.Tx()
	if tx == nil {
		return nil
	}

	var keyID int64
	err := tx.QueryRow("SELECT id FROM cache_key where key = ?", cacheKey).Scan(&keyID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("DELETE FROM cache_value WHERE key_id = ?", keyID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("DELETE FROM cache_key WHERE id = ?", keyID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
