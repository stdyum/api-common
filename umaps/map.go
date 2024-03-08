package umaps

func MapFuncErr[K1 comparable, V1 any, K2 comparable, V2 any](m map[K1]V1, converter func(key K1, value V1) (K2, V2, error)) (map[K2]V2, error) {
	newMap := make(map[K2]V2, len(m))
	for key, value := range m {
		newKey, newValue, err := converter(key, value)
		if err != nil {
			return nil, err
		}

		newMap[newKey] = newValue
	}
	return newMap, nil
}

func MapFunc[K1 comparable, V1 any, K2 comparable, V2 any](m map[K1]V1, converter func(key K1, value V1) (K2, V2)) map[K2]V2 {
	newMap, _ := MapFuncErr(m, func(key K1, value V1) (K2, V2, error) {
		newKey, newValue := converter(key, value)
		return newKey, newValue, nil
	})

	return newMap
}
