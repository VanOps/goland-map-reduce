package main

import "fmt"
import "github.com/VanOps/goland-map-reduce/functionMap"
import "github.com/VanOps/goland-map-reduce/functionReduce"

func main() {
	f := func(i int) int {
		return i * 2
	}

	// Define a function to be passed to ReduceIntToInt
	r := func(x, y int) int {
		return x + y
	}

	// Define a slice of integers
	a := []int{1, 2, 3, 4, 5}

	// Call MapIntToInt with the function and slice as arguments
	result := functionMap.MapIntToInt(f, a)

	fmt.Println(result)

	// Call ReduceIntToInt with the function and slice as arguments
	reducedResult := functionReduce.ReduceIntToInt(r, a)
	fmt.Println(reducedResult)
}
