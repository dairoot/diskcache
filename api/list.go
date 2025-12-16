package api

import (
	"fmt"
)

func (dc *DiskCache) LPush(cacheKey string, cacheValue string) error {
	tx := dc.Tx()
	if tx == nil {
		return nil
	}

	keyID, _ := GetKeyIDByCU(tx, cacheKey)
	valueMd5 := GetMd5String(cacheValue)

	// 插入内容
	_, err := tx.Exec("INSERT INTO cache_value(key_id, value, value_md5) VALUES(?,?,?)",
		keyID, cacheValue, valueMd5)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (dc *DiskCache) LPop(cacheKey string) (string, error) {
	return dc._pop(cacheKey, "left")
}
func (dc *DiskCache) RPop(cacheKey string) (string, error) {
	return dc._pop(cacheKey, "right")
}

func (dc *DiskCache) _pop(cacheKey string, turnTo string) (string, error) {
	orderBy := "id asc"

	if turnTo == "left" {
		orderBy = "id desc"
	}

	keyID, err := dc.GetKeyIDNotTx(cacheKey)

	if err != nil {
		return "", err
	}

	tx := dc.Tx()
	if tx == nil {
		return "", nil
	}

	var value string
	var valueID int64
	query := fmt.Sprintf("SELECT id, value FROM cache_value WHERE key_id = ? ORDER BY %s", orderBy)
	err = tx.QueryRow(query, keyID).Scan(&valueID, &value)

	if err != nil {
		tx.Rollback()
		return "", err
	}

	_, err = tx.Exec("DELETE FROM cache_value WHERE id = ?", valueID)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	tx.Commit()
	return value, nil
}

func (dc *DiskCache) RRange(cacheKey string, offset int64, limit int64) []string {
	return dc._range(cacheKey, offset, limit, "right")
}
func (dc *DiskCache) LRange(cacheKey string, offset int64, limit int64) []string {
	return dc._range(cacheKey, offset, limit, "left")
}

func (dc *DiskCache) _range(cacheKey string, offset int64, limit int64, turnTo string) []string {
	orderBy := "id asc"

	if turnTo == "left" {
		orderBy = "id desc"
	}
	keyID, err := dc.GetKeyIDNotTx(cacheKey)

	if err != nil {
		return []string{}
	}

	tx := dc.Tx()
	if tx == nil {
		return []string{}
	}
	defer tx.Commit()

	query := fmt.Sprintf("SELECT value FROM cache_value WHERE key_id = ? ORDER BY %s limit ?, ?", orderBy)
	rows, err := tx.Query(query, keyID, offset, limit)

	if err != nil {
		return []string{}
	}
	defer rows.Close()

	// 遍历结果
	var values []string
	for rows.Next() {
		var value string
		if err := rows.Scan(&value); err != nil {
			continue
		}
		values = append(values, value)
	}

	return values
}

func (dc *DiskCache) LLen(cacheKey string) int64 {
	keyID, err := dc.GetKeyIDNotTx(cacheKey)

	if err != nil {
		return 0
	}

	var c int64
	dc.Conn.QueryRowContext(dc.Ctx, "SELECT count(*) FROM cache_value WHERE key_id = ?", keyID).Scan(&c)
	return c
}

func (dc *DiskCache) LRem(cacheKey string, cacheValue string) error {
	keyID, err := dc.GetKeyIDNotTx(cacheKey)
	valueMd5 := GetMd5String(cacheValue)

	if err != nil {
		return err
	}

	tx := dc.Tx()
	if tx == nil {
		return nil
	}

	tx.Exec("DELETE FROM cache_value WHERE key_id = ? and value_md5 = ?;", keyID, valueMd5)
	return tx.Commit()
}
