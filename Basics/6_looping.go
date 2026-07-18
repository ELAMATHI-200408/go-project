// Six problems
// 1. Print this
// *
// *
// *
//package main
//import "fmt"
//func main(){
//for i := 0; i<3; i++{
//	fmt.Println("*")
//}

//}

// 2. Print this
// 1
// 2 2
// 3 3 3
// 4 4 4 4
//package main
//import "fmt"
//func main(){
//for i := 1; i <= 4; i++ {
//for j := 1; j <= i; j++ {
//	fmt.Print(i," ")
//}
//fmt.Println()
//}

//}

// 3. Print this
// *
// * *
// * * *
// * * * *
//package main
//import "fmt"
//func main(){
//for i := 1; i <= 4; i++ {
//for j := 1; j <= i; j++ {
//fmt.Print("*" ," ")
//}
// fmt.Println()
//}

//}

// 4. Print this.. if odd print 1, if even print 0
// 1
// 0 0
// 1 1 1
// 0 0 0 0
//package main
//import "fmt"
//func main(){
//for i := 1; i <= 4; i++ {
//for j := 1; j <= i; j++ {
//if (i%2==0){
//fmt.Print("0"," ")
//}else{
//fmt.Print("1"," ")
//}
// }
// fmt.Println()
//}

//}

// 5. Print this..print only odd rows / even rows
// 1
// 3 3 3
// 5 5 5 5 5

// or

// 2 2
// 4 4 4 4
//package main
//import "fmt"
//func main(){
//for i := 1; i <= 6; i++ {
//for j := 1; j <=i; j++ {
//if (i%2==0){
//continue
// }else{
//fmt.Print(i," ")
//}
//}
//fmt.Println()
//}

//}

// 6. Print this
// 1
// 2 3
// 4 5 6
// 7 8 9 10
package main

import "fmt"

func main() {
	num := 1
	for i := 1; i <= 4; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(num, " ")
			num++
		}
		fmt.Println()
	}

}
