package main

import (
	"fmt"

	textdistance "github.com/masatana/go-textdistance"
)

func main() {
	a := "Hello"
	b := "Help"
	fmt.Println(a, b, "diff", textdistance.LevenshteinDistance(a, b))
}
