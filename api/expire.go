package api

func (dc *DiskCache) Expire(cacheKey string, ttl float64) error {
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

	err = UpdateKeyID(tx, keyID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = UpdateKeyIDTTL(tx, keyID, ttl)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
