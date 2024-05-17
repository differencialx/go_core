package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))

	return len(bs), nil
}
