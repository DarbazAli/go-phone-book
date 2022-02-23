package main

import (
	"fmt"
	"os"
	"path"
)

type Entry struct {
	Name    string
	Surname string
	Tel     string
}

var data = []Entry{}

// search function
func search(key string) *Entry {
	for i, v := range data {
		if v.Surname == key {
			return &data[i]
		}
	}
	return nil
}

// list function
func list() {
	for _, v := range data {
		fmt.Println(v)
	}
}

// function main begins program execution
func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		exe := path.Base(arguments[0])
		fmt.Printf("Usage: %s search|list <arguments>\n", exe)
		return
	}

	// mimic the data
	data = append(data, Entry{"John", "Doe", "1234"})
	data = append(data, Entry{"Jane", "Doe", "1234"})
	data = append(data, Entry{"Black", "Doe", "1234"})

	// Defferentiate between teh commands
	switch arguments[1] {

	// the search command
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search Surname")
			return
		}
		result := search(arguments[2])
		if result == nil {
			fmt.Println("Entry not found:", arguments[2])
			return
		}
		fmt.Println(*result)

	// the list command
	case "list":
		list()

	default:
		fmt.Println("Not a valid option")

	}
}
