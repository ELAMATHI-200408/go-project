package main

import (
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print(string(content))
}
