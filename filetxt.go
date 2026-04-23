package main

import (
	"fmt"
	"os"
	"strings"
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
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".txt") {
			fmt.Println(entry.Name())
		}

	}

}
