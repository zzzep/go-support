package setter

// SetIfNotEmpty sets the value of target to value only if value is not empty.
// For strings, empty means empty string "".
// For slices, empty means nil or length 0.
// For maps, empty means nil or length 0.
// For pointers, empty means nil.
// For other types, empty means the zero value.
func SetIfNotEmpty[T comparable](target *T, value T) {
	if !isEmpty(value) {
		*target = value
	}
}

// SetIfNotEmptyString sets the value of target to value only if value is not an empty string.
func SetIfNotEmptyString(target *string, value string) {
	if value != "" {
		*target = value
	}
}

// SetIfNotEmptySlice sets the value of target to value only if value is not nil or empty.
func SetIfNotEmptySlice[T any](target *[]T, value []T) {
	if len(value) > 0 {
		*target = value
	}
}

// SetIfNotEmptyMap sets the value of target to value only if value is not nil or empty.
func SetIfNotEmptyMap[K comparable, V any](target *map[K]V, value map[K]V) {
	if len(value) > 0 {
		*target = value
	}
}

// isEmpty checks if a value is considered empty.
// For strings, empty means empty string "".
// For slices, empty means nil or length 0.
// For maps, empty means nil or length 0.
// For pointers, empty means nil.
// For other types, empty means the zero value.
func isEmpty[T comparable](v T) bool {
	var zero T
	return v == zero
}
