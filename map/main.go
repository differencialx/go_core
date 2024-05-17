package main

import "fmt"

func main() {
	colors := make(map[string]string)

	colors["red"] = "#010101"
	colors["green"] = "#010111"
	colors["white"] = "#ffffff"

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
