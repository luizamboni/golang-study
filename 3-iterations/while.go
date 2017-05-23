package main

import "fmt"

func main() {
  
  i := 0
  
  /* not exists while keyword in go */
  for i < 20 {
    /* in go ++ is a statemant, not an expression
       and can not stay in another expression
    */
    i++
    fmt.Println("i is: ",i)
  }
}