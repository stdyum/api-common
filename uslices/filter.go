package uslices

func FilterFunc[T any](arr []T, filter func(item T) bool) []T {
	newArr := make([]T, 0, len(arr))
	for _, item := range arr {
		f := filter(item)
		if !f {
			continue
		}

		newArr = append(newArr, item)
	}

	return newArr
}
