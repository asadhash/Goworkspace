package main

import "fmt"

func wrapper(alpha int) func() int {
	x := alpha
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper(5)
	fmt.Println(increment())
	fmt.Println(increment())
}
