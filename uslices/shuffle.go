package uslices

import (
	"math/rand"
)

func Shuffle[T any](array []T) []T {
	for i := len(array) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		array[i], array[j] = array[j], array[i]
	}

	return array
}
