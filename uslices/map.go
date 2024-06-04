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

func ToMapFuncErr[T any, K comparable, V any](arr []T, converter func(item T) (K, V, error)) (map[K]V, error) {
	m := make(map[K]V, len(arr))
	for _, item := range arr {
		key, value, err := converter(item)
		if err != nil {
			return nil, err
		}

		m[key] = value
	}

	return m, nil
}

func ToMapFunc[T any, K comparable, V any](arr []T, converter func(item T) (K, V)) map[K]V {
	m, _ := ToMapFuncErr(arr, func(item T) (K, V, error) {
		k, v := converter(item)
		return k, v, nil
	})
	return m
}

func MapMapFuncErr[IK comparable, IV any, K comparable, V any](mm map[IK]IV, converter func(key IK, value IV) (K, V, error)) (map[K]V, error) {
	m := make(map[K]V, len(mm))
	for k, v := range mm {
		key, value, err := converter(k, v)
		if err != nil {
			return nil, err
		}

		m[key] = value
	}

	return m, nil
}

func MapMapFunc[IK comparable, IV any, K comparable, V any](mm map[IK]IV, converter func(key IK, value IV) (K, V)) map[K]V {
	m, _ := MapMapFuncErr(mm, func(k IK, v IV) (K, V, error) {
		key, value := converter(k, v)
		return key, value, nil
	})
	return m
}
