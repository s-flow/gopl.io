package main

import (
	"bufio"
	"fmt"
	"gopl.io/ch2/tempconv"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) <= 1 {
		reader := bufio.NewScanner(os.Stdin)
		fmt.Print("> ")
		for {

		}
	}

	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf(" %s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}
