package main

import "fmt"

func main() {
	var numbers = [5]int64{2, 3, 4, 5, 5}
	var sum int64
	for i := len(numbers) - 1; i >= len(numbers)-3; i-- {
		fmt.Printf("Element %d is: %d\n", i, numbers[i])
		sum += numbers[i]
	}
	fmt.Printf("Sum is: %d", sum)
}
