package ex2_1

import (
	"fmt"
	"testing"
)

func TestKevinToCelius(t *testing.T) {
	c := KToC(0)
	fmt.Printf("%v\n", c)
	k := CToK(-273.15)
	fmt.Println(k)
}
