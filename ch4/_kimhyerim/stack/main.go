package main

import "fmt"

func main() {
	stack := []string{}

	// push
	stack = append(stack, "A")
	fmt.Println(stack)
	stack = append(stack, "B")
	fmt.Println(stack)
	stack = append(stack, "C")
	fmt.Println(stack)

	// pop
	stack = stack[:len(stack)-1]
	fmt.Println(stack)
	stack = stack[:len(stack)-1]
	fmt.Println(stack)
	stack = stack[:len(stack)-1]
	fmt.Println(stack)
}
