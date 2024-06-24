package uslices

func Distinct[T comparable](array []T) []T {
	uniqueMap := make(map[T]bool)
	for _, v := range array {
		uniqueMap[v] = true
	}

	distinctArray := make([]T, 0, len(uniqueMap))
	for key := range uniqueMap {
		distinctArray = append(distinctArray, key)
	}

	return distinctArray
}
