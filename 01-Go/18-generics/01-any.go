package generics

import (
	"fmt"
)

func Any() {
	PrintList(1, 2, 3, 4, 5)
	PrintList("a", "b", "c", "d", "e")
}

func PrintList(list ...any) {
	for _, v := range list {
		fmt.Println(v)
	}
}
