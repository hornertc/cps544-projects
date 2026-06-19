package main

import (
	"errors"
	"fmt"
)

type ErrorList []error

// func Error() string
func (el ErrorList) Error() (s string) {
	for _, e := range el {
		s += "-" + e.Error()
	}
	return
}

// implement errors.Is(errorList, specific) support
func (el ErrorList) Is(err error) bool {
	for _, e := range el {
		if errors.Is(e, err) {
			return true
		}
	}
	return false
}

func main() {
	errA := errors.New("A")
	errB := errors.New("B")
	errC := errors.New("C")

	errs := ErrorList{errA, errB}

	var myError error = errs

	fmt.Println("A", myError == errA, errors.Is(myError, errA))
	fmt.Println("B", myError == errB, errors.Is(myError, errB))
	fmt.Println("C", myError == errC, errors.Is(myError, errC))
}
