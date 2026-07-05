// Swap two variables without a temp variable using multiple assignment
package main
import "fmt"
func main(){
	//assign a  and b values
	a := 10
	b := 20
	//swap the a  and b values
	a,b = b,a
	//a = b
	//b = a
  fmt.Println(a)
  fmt.Println(b)

}