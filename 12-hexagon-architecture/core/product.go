package core

import "fmt"

type Product struct {
	Name     string
	Ean      string
	Category string
}

func (s Product) String() string {
	return fmt.Sprintf("<Name: '%v' Ean: '%v' Category: '%v' />", s.Name, s.Ean, s.Category)
}
