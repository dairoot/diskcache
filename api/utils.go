package api

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"time"
)

func GetMd5String(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	md5Hash := hash.Sum(nil)

	md5String := hex.EncodeToString(md5Hash)
	return md5String
}

func GetValue(tx *sql.Tx, keyID int64) (string, error) {
	var value string
	err := tx.QueryRow("SELECT value FROM cache_value where key_id = ?", keyID).Scan(&value)
	return value, err
}

func UpdateKeyIDTTL(tx *sql.Tx, keyID int64, ttl float64) error {

	var expireTime *int64 = nil
	if ttl != 0 {
		expire := time.Now().Add(time.Duration(ttl*1000) * time.Millisecond).Unix()
		expireTime = &expire
	}

	_, err := tx.Exec("UPDATE cache_key SET expire_time=? WHERE id = ?", expireTime, keyID)
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func UpdateKeyID(tx *sql.Tx, keyID int64) error {
	accessTime := time.Now().Unix()
	_, err := tx.Exec("UPDATE cache_key SET access_time = ?, access_count=access_count+1  WHERE id = ?", accessTime, keyID)
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func InsertKeyID(tx *sql.Tx, cacheKey string) (int64, error) {
	accessTime := time.Now().Unix()
	createTime := time.Now().Unix()

	result, err := tx.Exec("INSERT INTO cache_key(key,expire_time,access_time,create_time) VALUES(?,?,?,?)",
		cacheKey, nil, accessTime, createTime)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	keyID, err := result.LastInsertId()
	return keyID, err
}

func getKeyID(tx *sql.Tx, cacheKey string) (int64, bool) {
	var keyID int64
	var expireTime sql.NullFloat64
	var isExists bool

	err := tx.QueryRow("SELECT id, expire_time FROM cache_key where key = ?", cacheKey).Scan(&keyID, &expireTime)
	nowTime := time.Now().Unix()
	if err != nil {
		// 不存在
		isExists = false
	} else if !expireTime.Valid {
		// expire_time 为 NULL，表示永不过期
		isExists = true
	} else if expireTime.Float64 <= float64(nowTime) {
		// 已过期
		isExists = false
	} else {
		isExists = true
	}

	return keyID, isExists
}

func GetKeyIDByCU(tx *sql.Tx, cacheKey string) (int64, int16) {
	// 获取 keyID，没有则创建，有则更新
	var keyID int64
	keyID, isExists := getKeyID(tx, cacheKey)

	if keyID != 0 {
		UpdateKeyID(tx, keyID)
	} else {
		var err error
		keyID, err = InsertKeyID(tx, cacheKey)
		if err != nil {
			return 0, 1
		}
	}

	if isExists {
		return keyID, 0
	}
	return keyID, 1
}
