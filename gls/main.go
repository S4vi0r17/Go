package main

import (
	"flag"
	"fmt"
)

func main() {
	// Filter by pattern
	flagPattern := flag.String("p", "", "filter by pattern")
	flagAll := flag.Bool("a", false, "show all files")
	flagNumberRecords := flag.Int("n", 0, "number of records to show")

	// Order flags
	hasOrderByTime := flag.Bool("t", false, "order by time, oldest first")
	hasOrderBySize := flag.Bool("s", false, "order by size, smallest first")
	hasOrderReverse := flag.Bool("r", false, "reverse the order")

	flag.Parse() // parse the flags

	fmt.Println(*flagPattern)
	fmt.Println(*flagAll)
	fmt.Println(*flagNumberRecords)

	fmt.Println(*hasOrderByTime)
	fmt.Println(*hasOrderBySize)
	fmt.Println(*hasOrderReverse)
}
