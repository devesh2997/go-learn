package main

import (
	"fmt"
)

// variable declaration
// if else conditions, switch
// looping
// functions
// packages
// structs
// interface
// error handling
// go routines
// channels
// type inference

func main() {
	fmt.Println("Hello World")

	var temp int = 3
	fmt.Println(temp)

	temp1 := "a"
	fmt.Println(temp1)

	if temp > 2 {
		fmt.Println("something")
		fmt.Println("something")
	} else if temp < 1 {
		fmt.Println("do ")
	}

	switch temp1 {
	case "a":
		fmt.Println("is a")
		fallthrough
	case "b":
		fmt.Println("is b")
		fallthrough
	case "c":
	default:
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	var arr []int = []int{1, 2}
	fmt.Println(arr)

	arr1 := []int{}
	fmt.Println(arr1)
	arr1 = append(arr1, 2, 3, 4, 5, 5, 5)
	fmt.Println(arr1)

	for _, val := range arr1 {
		fmt.Println(val)
	}
	res, res1 := add(2, 3)
	fmt.Println(res)
	fmt.Println(res1)

	_ = addAll(3, 4, 5, 5, 6)
	_ = addAll(3, 2)
}

func add(a int, b int) (int, int) {
	return a + b, 0
}

func addAll(args ...int) int {
	sum := 0
	for _, arg := range args {
		sum += arg
	}

	return sum
}
