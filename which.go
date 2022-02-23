package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// function main begins program execution
// main also a goroutine
func main() {
	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Println("Please provide an argument")
		return
	}

	file := arguments[1]
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)

	for _, dir := range pathSplit {
		fullPath := filepath.Join(dir, file)

		// Does it exist?
		fileInfo, err := os.Stat(fullPath)

		if err == nil {
			mode := fileInfo.Mode()

			// Is it a regular file?
			if mode.IsRegular() {
				// Is it excutable?
				if mode&0111 != 0 {
					fmt.Println(fullPath)
					return
				}
			}
		}
	}
}
