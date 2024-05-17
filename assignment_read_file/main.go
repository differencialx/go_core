package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(string(fileBytes))
}
