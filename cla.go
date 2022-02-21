package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	// get data from cmd arguments
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need one more argument!")
		return
	}

	// processing phase
	var min, max float64

	for i := 1; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)

		if err != nil {
			continue
		}
		if n < min {
			min = n
		}

		if n > max {
			max = n
		}
	}

	fmt.Println("Max:", max)
	fmt.Println("Min:", min)

}
