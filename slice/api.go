package slice

type mappingFunctionDefinition[InType any, OutType any] func(item InType) OutType

func Map[InType any, OutType any](items []InType, mappingFunction mappingFunctionDefinition[InType, OutType]) []OutType {
	var result []OutType

	for _, item := range items {
		result = append(result, mappingFunction(item))
	}

	return result
}

type reduceFunctionDefinition[InType any, OutType any] func(item InType, sum OutType) OutType

func Reduce[InType any, OutType any](items []InType, reduceFunction reduceFunctionDefinition[InType, OutType]) OutType {
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

// https://go.dev/doc/go1.18
// The Go compiler cannot handle type declarations inside generic functions or methods.
//We hope to provide support for this feature in a future release.
//func (p *pipeline[T]) Map[F any](mapFunc func(T) F) *pipeline[F] {
func (p *pipeline[T]) Map(mapFunc mappingFunctionDefinition[T, T]) *pipeline[T] {
	return Pipeline(Map(p.items, mapFunc))
}

func (p *pipeline[T]) Reduce(reduceFunc reduceFunctionDefinition[T, T]) T {
	var result T
	result = Reduce(p.items, reduceFunc)
	return result
}

func (p *pipeline[T]) Collect() []T {
	return p.items
}
