package main

import (
	"fmt"
	"os"
)

func main() {
	var foldername string
	fmt.Scanln(&foldername)
	entries, err := os.ReadDir(foldername)

	if err != nil {
		fmt.Println(err)
		return
	}
	for _, entry := range entries {
		fmt.Println(entry.Name())

	}

}
