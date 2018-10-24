package util

func ContainsInt(array []int, find int) bool {
	for _, a := range array {
		if a == find {
			return true
		}
	}
	return false
}
