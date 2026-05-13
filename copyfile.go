package main

import (
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("example.txt")
	file, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	file.WriteString(string(content))
	fmt.Print(string(content))
}
