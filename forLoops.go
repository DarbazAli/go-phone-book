package main

import "fmt"

func main() {
	// Traditional for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i*i, " ")
	}
	// print a blank line
	fmt.Println()

	// For loop used as while loop
	i := 0
	for {
		if i == 10 {
			break
		}

		fmt.Println(i*i, " ")
		i++
	}
	fmt.Println()

	// Using range keyword
	arr := [5]int{1, 2, 3, 4, 5}
	for i, value := range arr {
		fmt.Println("Index:", i, "Value", value)
	}

}
