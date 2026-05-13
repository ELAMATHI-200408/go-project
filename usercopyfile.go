package main

import (
	"fmt"
	"os"
)

func main() {
	var sourcefile string
	fmt.Print("Enter source file: ")
	fmt.Scan(&sourcefile)
	content, err := os.ReadFile(sourcefile)
	var destinationfile string
	fmt.Print("Enter destination file: ")
	fmt.Scan(&destinationfile)
	file, err := os.Create(destinationfile)
	if err != nil {
		fmt.Println(err)
		return
	}
	file.WriteString(string(content))
	fmt.Println("Stored in:", destinationfile)
	fmt.Print(string(content))
}
