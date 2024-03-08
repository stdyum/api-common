package uslices

func MapFuncErr[T any, R any](arr []T, converter func(item T) (R, error)) ([]R, error) {
	newArr := make([]R, len(arr))
	for i, item := range arr {
		newItem, err := converter(item)
		if err != nil {
			return nil, err
		}

		newArr[i] = newItem
	}
	return newArr, nil
}

func MapFunc[T any, R any](arr []T, converter func(item T) R) []R {
	newArr, _ := MapFuncErr(arr, func(item T) (R, error) {
		return converter(item), nil
	})

	return newArr
}

func ToMapFunc[T any, K comparable, V any](arr []T, converter func(item T) (K, V)) map[K]V {
	m := make(map[K]V, len(arr))
	for _, item := range arr {
		key, value := converter(item)
		m[key] = value
	}

	return m
}
