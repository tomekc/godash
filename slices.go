package godash

func Map[E, T any](input []E, mapper func(E) T) []T {
	result := make([]T, len(input))
	for i, v := range input {
		result[i] = mapper(v)
	}
	return result
}

func Filter[E any](input []E, matcher func(E) bool) []E {
	result := make([]E, 0, len(input))
	for _, v := range input {
		if matcher(v) {
			result = append(result, v)
		}
	}
	return result
}

func ForEach[E any](input []E, iteratee func(E)) {
	for _, v := range input {
		iteratee(v)
	}
}

func Contains[E comparable](haystack []E, needle E) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}
