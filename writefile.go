package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("hello.txt")
	//file, err := os.OpenFile("hello.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	file.WriteString("Hi Buddy")
	file.Close()
}
