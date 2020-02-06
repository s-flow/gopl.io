package main

import "fmt"

func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}

func zero_new(ptr *[32]byte) {
	*ptr = [32]byte{}
}

func main() {
	ptr := [32]byte{1, 2, 3}
	fmt.Printf("zero before ptr: %x\n", ptr)
	zero(&ptr)

	fmt.Printf("zero after ptr: %x\n", ptr)

	ptr = [32]byte{4, 5, 6}
	fmt.Printf("zero before ptr: %x\n", ptr)
	zero_new(&ptr)

	fmt.Printf("zero_new after ptr: %x\n", ptr)
}
