package main

import "fmt"

func main() {

  // declare but not assing
  var a string
  fmt.Println("a:", a)
  a = "teste assign"
  fmt.Println("a:",a)

  // simple one variable with assign
  var b string = "teste"
  fmt.Println(b)

  // 2 variables with assign
  var c, d string = "bbb", "ccc"
  fmt.Println(c,d)

  // shot syntax
  e := "eeee"
  fmt.Println(e)

}