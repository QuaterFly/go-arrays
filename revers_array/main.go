package main

import "fmt"

func main() {
	var numbers = []int64{2, 3, 4, 5, 5}
	var reversNumbers []int64
	for i := len(numbers) - 1; i >= 0; i-- {
		reversNumbers = append(reversNumbers, numbers[i])
	}
	fmt.Printf("Current array is: %v\n", numbers)
	fmt.Printf("Revert array is: %v\n", reversNumbers)
}
