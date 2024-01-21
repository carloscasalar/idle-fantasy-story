package utils

// NoNilValue returns the value if it is not nil, otherwise it returns the default value
func NoNilValue[T any](value *T, defaultValue T) T {
	if value == nil {
		return defaultValue
	}
	return *value
}
