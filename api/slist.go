package api

func (dc *DiskCache) SAdd(cacheKey string, cacheValue string) error {

	tx := dc.Tx()

	defer tx.Commit()

	keyID, _ := GetKeyIDByCU(tx, cacheKey)

	vauleMd5 := GetMd5String(cacheValue)

	// 插入并更新内容
	var vauleID int

	err := tx.QueryRow("SELECT id FROM cache_value where key_id = ? and value_md5 = ?", keyID, vauleMd5).Scan(&vauleID)

	if err == nil {
		_, err := tx.Exec("UPDATE cache_value SET value = ? , value_md5 =? WHERE id = ?",
			cacheValue, vauleMd5, vauleID)
		if err != nil {
			tx.Rollback()
			return err
		}

	} else {
		_, err := tx.Exec("INSERT INTO cache_value(key_id,value,value_md5) VALUES(?,?, ?)", keyID, cacheValue, vauleMd5)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return nil
}

func (dc *DiskCache) SPop(cacheKey string) (string, error) {
	return dc.LPop(cacheKey)
}

func (dc *DiskCache) SRem(cacheKey string, cacheValue string) error {
	return dc.LRem(cacheKey, cacheValue)
}
