package api

import (
	"strconv"
)

func (dc *DiskCache) Incr(cacheKey string) int64 {

	tx := dc.Tx()

	defer tx.Commit()

	keyID, _ := GetKeyIDByCU(tx, cacheKey)

	// 插入并更新内容
	var vauleID int
	var value string

	err := tx.QueryRow("SELECT id, value FROM cache_value where key_id = ?", keyID).Scan(&vauleID, &value)

	if err == nil {
		_, err := tx.Exec("UPDATE cache_value SET value = value+1  WHERE id = ?", vauleID)

		if err != nil {
			tx.Rollback()
			return 0
		}

		result, _ := strconv.ParseInt(value, 10, 64)
		return result
	} else {

		_, err := tx.Exec("INSERT INTO cache_value(key_id,value) VALUES(?,?)", keyID, 1)
		if err != nil {
			tx.Rollback()
			return 0
		}
	}
	return 0

}
