package slice

func Map[InType any, OutType any](items []InType, mappingFunction func(item InType) OutType) []OutType {
	var result []OutType

	for _, item := range items {
		result = append(result, mappingFunction(item))
	}

	return result
}

func Reduce[InType any, OutType any](items []InType, reduceFunction func(item InType, sum OutType) OutType) OutType {
	var result OutType

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

func Filter[T any](items []T, filterFunction func(item T) bool) []T {
	var result []T

	for _, item := range items {
		if filterFunction(item) {
			result = append(result, item)
		}
	}

	return result
}
