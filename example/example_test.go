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

		sum := func(item int, result string) string {
			return result + fmt.Sprintf("%v", item)
		}

		result := slice.Reduce(numbers, sum)

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
})
