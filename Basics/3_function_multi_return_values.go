// Write a function that takes two ints and returns their sum and product (two return values)
package main
import "fmt"
//this func is used to calculate sum and product 
func calculate(a, b int) (int, int){
//declare sum
  sum := a+b
//declare product
  product := a*b
  //return sum and product
  return sum,product
} 
//main func
func main() {
   a := 10
   b := 20
  //call the calculate func
    sum, product := calculate(a, b)
//print sum and product
    fmt.Println("Sum:", sum)
   fmt.Println("Product:", product)
}






