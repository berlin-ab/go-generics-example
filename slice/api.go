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

type pipeline[T any] struct {
	items []T
}

func Pipeline[T any](items []T) *pipeline[T] {
	return &pipeline[T]{items: items}
}

func (p *pipeline[T]) Filter(filterFunc func(item T) bool) *pipeline[T] {
	return Pipeline[T](Filter(p.items, filterFunc))
}

func (p *pipeline[T]) Map(mapFunc func(item T) T) *pipeline[T] {
	return Pipeline(Map(p.items, mapFunc))
}

func (p *pipeline[T]) Collect() []T {
	return p.items
}
