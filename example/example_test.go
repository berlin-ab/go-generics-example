package example_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"go-generics-example/slice"
)

var _ = Describe("Example Go Generics", func() {
	It("can map", func() {
		numbers := []int{1, 2, 3, 4}

		double := func(item int) int {
			return item + item
		}

		doubled := slice.Map(numbers, double)

		Expect(doubled).To(Equal([]int{2, 4, 6, 8}))
	})

	It("can map to a different type", func() {
		numbers := []int{1, 2, 3, 4}

		stringify := func(item int) string {
			return fmt.Sprintf("%v", item)
		}

		strings := slice.Map(numbers, stringify)

		Expect(strings).To(Equal([]string{"1", "2", "3", "4"}))
	})

	It("can reduce", func() {
		numbers := []int{1, 2, 3, 4}

		sum := func(item int, total int) int {
			return total + item
		}

		result := slice.Reduce(numbers, sum)

		Expect(result).To(Equal(10))
	})

	It("can reduce to a different type", func() {
		numbers := []int{1, 2, 3, 4}

		append := func(item int, result string) string {
			return result + fmt.Sprintf("%v", item)
		}

		result := slice.Reduce(numbers, append)

		Expect(result).To(Equal("1234"))
	})

	It("can each", func() {
		numbers := []int{1, 2, 3, 4}
		added := []int{}

		slice.Each(numbers, func(item int) {
			increasedByOne := item + 1
			added = append(added, increasedByOne)
		})

		Expect(added).To(Equal([]int{2, 3, 4, 5}))
	})

	It("can filter", func() {
		numbers := []int{1, 2, 3, 4}

		result := slice.Filter(numbers, func(number int) bool {
			return number > 2
		})

		Expect(result).To(Equal([]int{3, 4}))
	})

	It("can pipeline without library help", func() {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		selectEvens := func(item int) bool {
			return (item % 2) == 0
		}

		double := func(item int) int {
			return item * 2
		}

		// The order of operations reads deepest to top-most
		result := slice.Map(
			slice.Filter(numbers, selectEvens),
			double)

		Expect(result).To(Equal([]int{4, 8, 12, 16, 20}))
	})

	It("can pipeline with same types", func() {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		selectEvens := func(item int) bool {
			return (item % 2) == 0
		}

		double := func(item int) int {
			return item * 2
		}

		result := slice.Pipeline(numbers).
			Filter(selectEvens).
			Map(double).
			Collect()

		Expect(result).To(Equal([]int{4, 8, 12, 16, 20}))
	})

	It("can reduce a pipeline with same types", func() {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		selectEvens := func(item int) bool {
			return (item % 2) == 0
		}

		double := func(item int) int {
			return item * 2
		}

		result := slice.Pipeline(numbers).
			Filter(selectEvens).
			Map(double).
			Reduce(func(item int, sum int) int {
				return sum + item
			})

		Expect(result).To(Equal(60))
	})

	//It("can pipeline with changing types", func() {
	//	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//
	//	selectEvens := func(item int) bool {
	//		return (item % 2) == 0
	//	}
	//
	//	double := func(item int) int {
	//		return item * 2
	//	}
	//
	//	stringify := func(item int) string {
	//		return ""
	//	}
	//
	//	result := slice.Pipeline(numbers).
	//		Filter(selectEvens).
	//		Map(double).
	//		Map(stringify).
	//		Collect()
	//
	//	Expect(result).To(Equal([]string{"4", "8", "12", "16", "20"}))
	//})
})
