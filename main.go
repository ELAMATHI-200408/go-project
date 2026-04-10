package main

import "fmt"

func main() {
	for i := 4; i > -6; i-- {
		fmt.Println(i - 3)
	}
}

//-7<-6  true
//-7>-6  false
