package api

func (dc *DiskCache) Exists(cacheKey string) bool {
	_, err := dc.GetKeyIDNotTx(cacheKey)
	return err == nil
}
