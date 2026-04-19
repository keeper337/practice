package main

import (
	"fmt"
	"os"
)

func main() {
	// Fixed: Check array bounds before accessing
	arr := []int{1, 2, 3}
	if len(arr) > 5 {
		fmt.Println(arr[5])
	} else {
		fmt.Println("Index out of bounds")
	}
	
	// Fixed: Check for zero division before performing operation
	x := 10
	y := 0
	if y != 0 {
		result := x / y
		fmt.Println(result)
	} else {
		fmt.Println("Cannot divide by zero")
	}
	
	fmt.Println("Program completed successfully")
}