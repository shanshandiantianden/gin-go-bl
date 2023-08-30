package main

import "fmt"

func main() {
	s := "hesslslo"
	switch {
	case s == "hello":
		fmt.Println("hello")
		fallthrough
	case s == "xxxx":
		fmt.Println("xxxx")
		fallthrough
	default:
		fmt.Println("world")

	}
}
