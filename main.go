package pslices

func Filter[T any](s []T, f func(T) bool) []T {
	newSlice := []T{}
	for _, item := range s {
		if f(item) {
			newSlice = append(newSlice, item)
		}
	}
	return newSlice
}

func Map[T any, V any](s []T, f func(T) V) []V {
	newSlice := make([]V, len(s))
	for i, item := range s {
		newSlice[i] = f(item)
	}
	return newSlice
}

func Contains[T comparable](s []T, v T) (int, bool) {
	for i, item := range s {
		if item == v {
			return i, true
		}
	}
	return -1, false
}
