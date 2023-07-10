package main

import "fmt"

func main() {

	silce := make([]int, 3, 3) // make([]T, length, capacity)

	silce[0] = 1
	silce[1] = 2
	silce[2] = 3

	silce = append(silce, 4)
	silce = append(silce, 5)
	silce = append(silce, 6)
	fmt.Println(silce)
	fmt.Println(len(silce))
	fmt.Println(cap(silce))
	silce = append(silce, 7)

	fmt.Println(silce)
	fmt.Println(len(silce))
	fmt.Println(cap(silce))

}
