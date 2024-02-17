package main

import (
	"errors"
	"fmt"
)

// define Error struct
type CustomError struct {
	msg string
}

// Create a function Error() string and associate it to the struct.
func (error *CustomError) Error() string {
	return error.msg
}

func returnError() error {
	return &CustomError{msg: "custom error"}
}

func main() {
	err := returnError()
	fmt.Println(errors.Unwrap(err))
	fmt.Println(err)
}

func Unwrap(err error) {
	panic("unimplemented")
}
