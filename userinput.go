package main

import (
	"fmt"
	"os"
)

func main() {
	var filename string
	//TODO: learn why & used here .. any other places in go we need to use &
	fmt.Scan(&filename)
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))

}
