package slice

func Map[T any](items []T, mappingFunction func(item T) T) []T {
	var result []T

	for _, item := range items {
		result = append(result, mappingFunction(item))
	}

	return result
}

func Reduce[T any](items []T, reduceFunction func(item T, sum T) T) T {
	var result T

	for _, item := range items {
		result = reduceFunction(item, result)
	}

	return result
}

func Each[T any](items []T, eachFunction func(item T)) {
	for _, item := range items {
		eachFunction(item)
	}
}
