package main

import(
	    "fmt"
        "os"
)

func main(){
	content, err := os.ReadFile("ascii.txt")
    file, err := os.Create("sample.txt")
  if err != nil {
	 fmt.Println(err)
	 return
}
	result := ""

  for _, char := range content {
		//fmt.Printf("Index %d: %c\n", i, r)
        //fmt.Printf("%c\n",r)
    if char == '&' || char == '$' || char == '*' {
        //fmt.Println("ignore it")
        continue 
        
    }
      result += string(char)//else {
        //concatenate the character to a string
        //fmt.Printf("%d %c",i,r)
    }
  
  //print the final string 
file.WriteString(result)

fmt.Println(result)
}
//string replace special char
//what is a regex