package api

func (dc *DiskCache) SAdd(cacheKey string, cacheValue string) error {
	tx := dc.Tx()
	if tx == nil {
		return nil
	}

	keyID, _ := GetKeyIDByCU(tx, cacheKey)
	valueMd5 := GetMd5String(cacheValue)

	// 插入并更新内容
	var valueID int

	err := tx.QueryRow("SELECT id FROM cache_value where key_id = ? and value_md5 = ?", keyID, valueMd5).Scan(&valueID)

	if err == nil {
		_, err := tx.Exec("UPDATE cache_value SET value = ?, value_md5 = ? WHERE id = ?",
			cacheValue, valueMd5, valueID)
		if err != nil {
			tx.Rollback()
			return err
		}
	} else {
		_, err := tx.Exec("INSERT INTO cache_value(key_id,value,value_md5) VALUES(?,?,?)", keyID, cacheValue, valueMd5)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (dc *DiskCache) SPop(cacheKey string) (string, error) {
	return dc.LPop(cacheKey)
}

func (dc *DiskCache) SRem(cacheKey string, cacheValue string) error {
	return dc.LRem(cacheKey, cacheValue)
}
