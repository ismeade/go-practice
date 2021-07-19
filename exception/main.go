// Package main provides ...
package main

import (
	"fmt"
)

func makeException() {
	fmt.Println("start makeException")
	panic("exception a")
	fmt.Println("end makeException")
}

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("recover:", err)
		}
	}()
	makeException()
	fmt.Println("main end")
}
