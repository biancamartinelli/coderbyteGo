package main
import (
  "fmt"
  "unicode"
)
func CodelandUsernameValidation(str string) string {

  l := len(str)

  if l < 4 || l > 25 {
    return "false"
  }

  if !unicode.IsLetter(rune(str[0])) {
    return "false"
  }

  if string(str[l-1])  == "_" {
    return "false"
  }

  for _, c := range str {
    if unicode.IsLetter(c) || unicode.IsDigit(c) || string(c) == "_" {
      continue
    } else {
      return "false"
    }
   }
  
  // code goes here  
  return "true";

}

func main() {

  // do not modify below here, readline is our function
  // that properly reads in the input for you
  fmt.Println(CodelandUsernameValidation(readline()))

}