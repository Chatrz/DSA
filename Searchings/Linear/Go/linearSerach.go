package linearsearch

func Search(key int, arr []int) int {
	for k, v := range arr {
		if v == key {
			return k
		}
	}
	return -1
}
