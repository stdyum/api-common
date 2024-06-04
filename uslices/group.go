package uslices

func GroupByMap[T any, K comparable](arr []T, getKey func(item T) K) map[K][]T {
	unique := make(map[K][]T)
	for _, item := range arr {
		key := getKey(item)
		unique[key] = append(unique[key], item)
	}
	return unique
}

func GroupBy[T any, K comparable](arr []T, getKey func(item T) K) [][]T {
	m := GroupByMap(arr, getKey)

	a := make([][]T, 0, len(m))

	for _, v := range m {
		a = append(a, v)
	}

	return a
}
