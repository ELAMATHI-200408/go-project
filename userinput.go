package main

import (
	"fmt"
	"os"
)

func main() {
	var filename string
	fmt.Scan(&filename)
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))

}
