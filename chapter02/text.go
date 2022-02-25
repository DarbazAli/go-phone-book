package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// clear cmd
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()

	// Strings
	aString := "Hello World! £"
	fmt.Println("First Character", string(aString[0]))

	// Runes
	r := '£'
	fmt.Println("As an int32 value:", r)
}
