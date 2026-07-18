// Check if a number is even/odd, positive/negative/zero using `switch`
// Get the number from user input
package main
import "fmt"
func main(){
	//var is keyword it declare which datatype is this
	var number int
	//get the num from user 
	fmt.Scanln(&number)
	//if else condition to check if the num is even or odd
	if number % 2 == 0 {
		fmt.Println("Even")
	}else{
		fmt.Println("Odd")
	}
	//switch case to print if the given num is (+|-|0)
	switch {
	//case condition
	case number>0: 
		  fmt.Println("Positive")
	case number<0:
		  fmt.Println("Negative")
	default:
		fmt.Println("Zero")
		
		
	}
}