package api

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

// LPush 将一个值插入到列表头部
func (dc *DiskCache) LPush(key string, value string) error {
	dc.mutex.Lock()
	defer dc.mutex.Unlock()

	valueList := []string{}

	keyInfo, err := dc.getKeyInfo(key)
	if err == nil {
		valueData, err := dc.getValueByKeyInfo(keyInfo)
		if err == nil {
			json.Unmarshal(valueData, &valueList)
		}
		// 删除旧的value文件
		dc.delValueFile(keyInfo.ValueHash)

	}
	valueList = append([]string{value}, valueList...)

	// 序列化整个列表
	valueListStr, err := json.Marshal(valueList)
	if err != nil {
		return err
	}

	// 保存列表数据到value文件
	valueDirPath, valueFileName, valueHash := dc.getValuePath(key, valueListStr)
	valueFilePath := filepath.Join(valueDirPath, valueFileName)

	if err := os.WriteFile(valueFilePath, valueListStr, 0644); err != nil {
		return err
	}

	// 更新key文件
	item := CacheItem{
		Key:       key,
		Time:      time.Now().Unix(),
		TTL:       0,
		ValueHash: valueHash,
	}

	data, err := json.Marshal(item)
	if err != nil {
		return err
	}

	dirPath, fileName := dc.getKeyPath(key)
	return os.WriteFile(filepath.Join(dirPath, fileName), data, 0644)
}

// LPop 移除并返回列表的第一个元素
func (dc *DiskCache) LPop(key string) (string, error) {
	return dc.listPop(key, "left")
}

// RPop 移除并返回列表的最后一个元素
func (dc *DiskCache) RPop(key string) (string, error) {
	return dc.listPop(key, "right")
}

func (dc *DiskCache) listPop(key string, turnTo string) (string, error) {
	dc.mutex.Lock()
	defer dc.mutex.Unlock()
	valueList := []string{}

	keyInfo, err := dc.getKeyInfo(key)
	if err != nil {
		return "", err
	}

	valueData, err := dc.getValueByKeyInfo(keyInfo)
	if err != nil {
		return "", err
	}

	if err := json.Unmarshal(valueData, &valueList); err != nil {
		return "", err
	}

	var fValue string
	var newValueList []string
	if turnTo == "left" {
		// 获取并移除第一个元素
		fValue = valueList[0]
		newValueList = valueList[1:]
	} else {
		// 获取并移除最后一个元素
		fValue = valueList[len(valueList)-1]
		newValueList = valueList[:len(valueList)-1]
	}

	if len(newValueList) == 0 {
		// 如果列表为空，删除键值对
		dc.delKeyFile(key)
		dc.delValueFile(keyInfo.ValueHash)
		return string(fValue), nil
	}

	// 更新列表
	updatedData, err := json.Marshal(newValueList)
	if err != nil {
		return "", err
	}

	// 计算新的value hash并保存
	valueDirPath, valueFileName, valueHash := dc.getValuePath(key, updatedData)
	valueFilePath := filepath.Join(valueDirPath, valueFileName)

	if err := os.WriteFile(valueFilePath, updatedData, 0644); err != nil {
		return "", err
	}

	// 删除旧的value文件
	dc.delValueFile(keyInfo.ValueHash)

	// 更新key文件
	keyInfo.ValueHash = valueHash
	keyInfo.Time = time.Now().Unix()

	keyData, err := json.Marshal(keyInfo)
	if err != nil {
		return "", err
	}

	dirPath, fileName := dc.getKeyPath(key)
	if err := os.WriteFile(filepath.Join(dirPath, fileName), keyData, 0644); err != nil {
		return "", err
	}

	return string(fValue), nil
}
