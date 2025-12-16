package api

import (
	"math/rand"
	"time"
)

func (dc *DiskCache) GetKeyIDNotTx(cacheKey string) (int64, error) {
	var keyID int64
	nowTime := time.Now().Unix()
	err := dc.Conn.QueryRowContext(dc.Ctx, "SELECT id FROM cache_key where key = ? AND (expire_time IS NULL OR expire_time > ?)",
		cacheKey, nowTime).Scan(&keyID)

	if err != nil {
		return 0, err
	}
	return keyID, nil
}

func (dc *DiskCache) Get(cacheKey string) (string, error) {
	keyID, err := dc.GetKeyIDNotTx(cacheKey)

	if err != nil {
		return "", err
	}

	var value string
	err = dc.Conn.QueryRowContext(dc.Ctx, "SELECT value FROM cache_value where key_id = ?", keyID).Scan(&value)

	if err != nil {
		return "", err
	}

	// 随机清理过期 key
	if rand.Intn(10) <= 1 {
		_ = dc.DelExpire()
	}
	return value, nil
}
