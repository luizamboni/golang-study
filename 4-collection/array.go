package main

import "fmt"

func main() {
  
  var a [5]int
  fmt.Println("emp:", a)

  a[4] = 100
  // access
  fmt.Println("emp:", a[4])
  
  fmt.Println("len:", len(a))

  // b, init values
  b := [5]int{1,2,4,6,12}
  fmt.Println("emp:", b)

}