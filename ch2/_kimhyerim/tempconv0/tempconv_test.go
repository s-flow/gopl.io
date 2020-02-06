// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package tempconv

import (
	"fmt"
	"testing"
)

func TestExample_one(t *testing.T) {
	{
		fmt.Printf("%g\n", BoilingC-FreezingC) // "100" °C
		boilingF := CToF(BoilingC)
		fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // "180" °F
	}
	// fmt.Printf("%g\n", boilingF-FreezingC)       // compile error: type mismatch

	// Output:
	// 100
	// 180
}

func TestExample_two(t *testing.T) {
	c := FToC(212.0)
	fmt.Println(c.String()) // "100°C"
	fmt.Printf("%v\n", c)   // "100°C"; no need to call String explicitly
	fmt.Printf("%s\n", c)   // "100°C"
	fmt.Println(c)          // "100°C"
	fmt.Printf("%g\n", c)   // "100"; does not call String
	fmt.Println(float64(c)) // "100"; does not call String

	// Output:
	// 100°C
	// 100°C
	// 100°C
	// 100°C
	// 100
	// 100
}
