package api

func (dc *DiskCache) SetNx(cacheKey string, cacheValue string, ttl float64) (int16, error) {
	tx := dc.Tx()

	defer tx.Commit()

	keyID, isInsert := GetKeyIDByCU(tx, cacheKey)
	UpdateKeyIDTTL(tx, keyID, ttl)

	// 插入并更新内容
	var vauleID int

	err := tx.QueryRow("SELECT id FROM cache_value where key_id = ?", keyID).Scan(&vauleID)
	if err == nil {
		_, err := tx.Exec("UPDATE cache_value SET value = ?  WHERE id = ?",
			cacheValue, vauleID)
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

	return isInsert, nil

}

func (dc *DiskCache) Set(cacheKey string, cacheValue string, ttl float64) error {
	_, err := dc.SetNx(cacheKey, cacheValue, ttl)
	return err
}
