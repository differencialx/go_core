package main

import (
	"fmt"
)

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, number := range numbers {
		reminder := number % 2
		if reminder == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}
}
