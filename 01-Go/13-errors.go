package main

import (
	"errors"
	"fmt"
	"strconv"
)

var errNotFound = errors.New("Not found error")

var food = map[int]string{
	1: "🍔",
	2: "🍕",
}

func Errors() {
	found, err := search("34")
	// if err == errNotFound {
	if errors.Is(err, errNotFound) {
		fmt.Println("Not found error occurred")
		return
	}
	if err != nil {
		fmt.Println("search: ", err)
		return
	}
	fmt.Println(found)
}

func search(key string) (string, error) {
	num, err := strconv.Atoi(key)
	if err != nil {
		// return "", err
		return "", fmt.Errorf("strconv.Atoi: %w", err)
	}

	f, err := findFood(num)
	if err != nil {
		// return "", err
		return "", fmt.Errorf("findFood: %w", err)
	}

	return f, nil

}

func findFood(id int) (string, error) {
	f, ok := food[id]
	if !ok {
		return "", errNotFound
	}
	return f, nil
}
