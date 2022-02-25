package main

import (
	"fmt"
	"os"
	"os/exec"
)

// function main begins program execution
func main() {

	// clear cmd on first run
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()

	// working with complex numbers
	c1 := 12 + 1i
	c2 := complex(5, 7)
	fmt.Printf("Type of c2: %T\n", c2)
	fmt.Printf("Type of c1: %T\n", c1)

	var c3 complex64 = complex64(c1 + c2)
	fmt.Println("c3:", c3)

	fmt.Printf("Type of c3: %T\n", c3)

	cZero := c3 - c3
	fmt.Println("cZero:", cZero)
	fmt.Println()

	// ===========================
	// Working with Integers
	// ===========================
	x := 12
	k := 5
	fmt.Println(x)
	fmt.Printf("Type of x: %T\n", x)

	div := x / k
	fmt.Println("div", div)

}
