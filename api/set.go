package api

func (dc *DiskCache) SetNx(cacheKey string, cacheValue string, ttl float64) (int16, error) {
	tx := dc.Tx()
	if tx == nil {
		return 0, nil
	}

	keyID, isInsert := GetKeyIDByCU(tx, cacheKey)
	if err := UpdateKeyIDTTL(tx, keyID, ttl); err != nil {
		return isInsert, err
	}

	// 插入并更新内容
	var valueID int

	err := tx.QueryRow("SELECT id FROM cache_value where key_id = ?", keyID).Scan(&valueID)
	if err == nil {
		_, err := tx.Exec("UPDATE cache_value SET value = ? WHERE id = ?",
			cacheValue, valueID)
		if err != nil {
			tx.Rollback()
			return isInsert, err
		}
	} else {
		_, err := tx.Exec("INSERT INTO cache_value(key_id,value) VALUES(?,?)", keyID, cacheValue)
		if err != nil {
			tx.Rollback()
			return isInsert, err
		}
	}

	tx.Commit()
	return isInsert, nil
}

func (dc *DiskCache) Set(cacheKey string, cacheValue string, ttl float64) error {
	_, err := dc.SetNx(cacheKey, cacheValue, ttl)
	return err
}
