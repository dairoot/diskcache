package api

import (
	"fmt"
	"log"
)

func (dc *DiskCache) LPush(cacheKey string, cacheValue string) error {

	tx := dc.Tx()

	defer tx.Commit()

	keyID, _ := GetKeyIDByCU(tx, cacheKey)

	// 插入内容
	_, err := tx.Exec("INSERT INTO cache_value(key_id, value) VALUES(?,?)",
		keyID, cacheValue)

	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (dc *DiskCache) LPop(cacheKey string) (string, error) {
	return dc.pop(cacheKey, "left")
}
func (dc *DiskCache) RPop(cacheKey string) (string, error) {
	return dc.pop(cacheKey, "right")
}

func (dc *DiskCache) pop(cacheKey string, turnTo string) (string, error) {
	orderBy := "id asc"

	if turnTo == "left" {
		orderBy = "id desc"
	}

	keyID, err := dc.GetKeyIDNotTx(cacheKey)

	if err != nil {
		return "", err
	}

	tx := dc.Tx()
	defer tx.Commit()

	var value string
	var valueID int64
	query := fmt.Sprintf("SELECT id, value FROM cache_value WHERE key_id = ? ORDER BY %s", orderBy)
	err = tx.QueryRow(query, keyID).Scan(&valueID, &value)

	if err != nil {
		return "", err
	}
	_, err = tx.Exec("DELETE FROM cache_value WHERE id = ?", valueID)
	if err != nil {
		return "", err
	}

	return value, nil

}

func (dc *DiskCache) RRange(cacheKey string, start int64, stop int64) []string {
	return dc.listRange(cacheKey, start, stop, "right")
}
func (dc *DiskCache) LRange(cacheKey string, start int64, stop int64) []string {
	return dc.listRange(cacheKey, start, stop, "left")
}

func (dc *DiskCache) listRange(cacheKey string, start int64, stop int64, turnTo string) []string {
	orderBy := "id asc"

	if turnTo == "left" {
		orderBy = "id desc"
	}
	keyID, err := dc.GetKeyIDNotTx(cacheKey)

	if err != nil {
		return []string{}
	}

	tx := dc.Tx()
	defer tx.Commit()
	query := fmt.Sprintf("SELECT value FROM cache_value WHERE key_id = ? ORDER BY %s limit ?, ?", orderBy)
	rows, err := tx.Query(query, keyID, start, stop)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// 遍历结果
	var values []string
	for rows.Next() {
		var value string
		if err := rows.Scan(&value); err != nil {
			log.Fatal(err)
		}
		values = append(values, value)

	}

	return values
}
