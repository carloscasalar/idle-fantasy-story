package utils

// PointerTo returns a pointer to the value
func PointerTo[T any](value T) *T {
	return &value
}
