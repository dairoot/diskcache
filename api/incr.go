package api

import (
	"strconv"
)

func (dc *DiskCache) Incr(cacheKey string) int64 {
	tx := dc.Tx()
	if tx == nil {
		return 0
	}

	keyID, _ := GetKeyIDByCU(tx, cacheKey)

	// 插入并更新内容
	var valueID int
	var value string

	err := tx.QueryRow("SELECT id, value FROM cache_value where key_id = ?", keyID).Scan(&valueID, &value)

	if err == nil {
		// 已存在，自增
		currentValue, _ := strconv.ParseInt(value, 10, 64)
		newValue := currentValue + 1

		_, err := tx.Exec("UPDATE cache_value SET value = ? WHERE id = ?", newValue, valueID)
		if err != nil {
			tx.Rollback()
			return 0
		}

		tx.Commit()
		return newValue
	}

	// 不存在，插入初始值 1
	_, err = tx.Exec("INSERT INTO cache_value(key_id,value) VALUES(?,?)", keyID, 1)
	if err != nil {
		tx.Rollback()
		return 0
	}

	tx.Commit()
	return 1
}
