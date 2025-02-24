package api

func (dc *DiskCache) Expire(cacheKey string, ttl float64) error {
	tx := dc.Tx()

	defer tx.Commit()

	var keyID int64
	err1 := tx.QueryRow("SELECT id FROM cache_key where key = ?", cacheKey).Scan(&keyID)
	err2 := UpdateKeyID(tx, keyID)
	err3 := UpdateKeyIDTTL(tx, keyID, ttl)

	if err1 != nil {
		return err1
	} else if err2 != nil {
		return err2
	} else if err3 != nil {
		return err3
	}
	return nil

}
