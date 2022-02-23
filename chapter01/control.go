package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	// make sure the user provides a cmd argument
	if len(os.Args) != 2 {
		fmt.Println("Please provide a command line argument")
		return
	}

	// grab the firt argument
	argument := os.Args[1]

	switch argument {
	case "0":
		fmt.Println("Zero")
	case "1":
		fmt.Println("One")
	case "2", "3", "4":
		fmt.Println("2 or 3 or 4")
		fallthrough // continue to next branch/case
	default:
		fmt.Println("Value: ", argument)

	}

	_, err := strconv.Atoi(argument)
	if err != nil {
		fmt.Println("Cannot convert to int: ", argument)
		return
	}

}
