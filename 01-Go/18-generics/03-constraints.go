package generics

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func Constraints() {
	fmt.Println(Includes([]int{1, 2, 3}, 2))
	fmt.Println(Includes([]string{"a", "b", "c"}, "d"))

	fmt.Println(Filter([]int{1, 2, 3, 4, 5}, func(n int) bool {
		return n > 3
	}))
}

func Includes[T comparable](list []T, value T) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

func Filter[T constraints.Ordered](list []T, filter func(T) bool) []T {
	var result []T
	for _, item := range list {
		if filter(item) {
			result = append(result, item)
		}
	}
	return result
}
