package util

// B must be a superclass of X.
func ArrayTransform[B any, X any](input []B) []X {
	result := make([]X, 0)
	for _, item := range input {
		result = append(result, any(item).(X))
	}

	return result
}
