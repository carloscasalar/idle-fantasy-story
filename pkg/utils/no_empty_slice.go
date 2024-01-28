package utils

func NoEmptySlice[T any](collection, defaultCollection []T) []T {
	if len(collection) == 0 {
		return defaultCollection
	}
	return collection
}
