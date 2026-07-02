package main
import(
	    "fmt"
	    "os"
      //if we use the Regexp we must include the regexp package
	    "regexp"
)
//the func starts here
func main(){
  //read the "workspace" file and that stored in content
	content, err := os.ReadFile("workspace.txt")
  //this portion is used in all programs for error handling
     if err != nil {
	 fmt.Println(err)
	 return
    }
   
  //this line is for used to remove numbers in paragraph
  //re := regexp.MustCompile("[0-9]") 
  //this line is for used to remove special characters in paragraph
  //re:=regexp.MustCompile("[^a-zA-z0-9]") 
  //this line is for used to remove special characters and numbers //re is an regexp object regexp is mention the package(^ means atha thavira)
  re:=regexp.MustCompile("[^a-zA-Z]")
  //this line replaces the all string ,stores in result and also it convert the byte to string using(string(content))
  result := re.ReplaceAllString(string(content),"")
  //print the output in terminal
  fmt.Println(result)
}