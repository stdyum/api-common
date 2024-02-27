package uslices

func ChunkFunc[T any](slice []T, chunkSize int, f func(chunk []T, i int)) {
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		f(slice[i:end], i)
	}
}
