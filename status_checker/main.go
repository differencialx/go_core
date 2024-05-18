package main

import (
	"fmt"
	"net/http"
)

func main() {
	c := make(chan string)
	c <- "Hi there!"
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
