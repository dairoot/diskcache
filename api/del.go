package api

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func listToString(list []int64) string {
	var newList []string
	for _, id := range list {
		newList = append(newList, strconv.FormatInt(id, 10))
	}
	return strings.Join(newList, ",")
}

func (dc *DiskCache) DelExpire() {
	nowTime := time.Now().Unix()
	rows, err := dc.Conn.QueryContext(dc.Ctx, "SELECT id FROM cache_key WHERE expire_time <= ? order by access_time asc limit 100", nowTime)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	// 遍历结果
	var keyIDS []int64
	for rows.Next() {
		var keyID int64
		if err := rows.Scan(&keyID); err != nil {
			log.Fatal(err)
		}
		keyIDS = append(keyIDS, keyID)
	}

	tx := dc.Tx()

	defer tx.Commit()

	sql1 := fmt.Sprintf("DELETE FROM cache_key WHERE id IN (%s);", listToString(keyIDS))
	sql2 := fmt.Sprintf("DELETE FROM cache_value WHERE key_id IN (%s);", listToString(keyIDS))
	tx.Exec(sql2)
	tx.Exec(sql1)

}

func (dc *DiskCache) Del(cacheKey string) error {

	tx := dc.Tx()

	defer tx.Commit()
	var keyID int64
	err := dc.Conn.QueryRowContext(dc.Ctx, "SELECT id FROM cache_key where key = ?", cacheKey).Scan(&keyID)

	if err != nil {
		return err
	}

	_, err2 := tx.Exec("DELETE FROM cache_key WHERE id = ?", keyID)
	_, err1 := tx.Exec("DELETE FROM cache_value WHERE key_id = ?", keyID)

	if err1 != nil {
		return err1
	}
	return err2
}
