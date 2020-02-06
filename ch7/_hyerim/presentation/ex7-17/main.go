package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack [][]string
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			elem := []string{tok.Name.Local}
			for _, attr := range tok.Attr {
				elem = append(elem, attr.Value)
			}
			stack = append(stack, elem) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s : %s\n", strings.Join(flatten(stack), " "), tok)
			}
		}
	}
}

func containsAll(x [][]string, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}

		if matches(x[0], y[0]) {
			y = y[1:]
		}

		x = x[1:]
	}
	return false
}

func matches(x []string, y string) bool {
	for _, elem := range x {
		if elem == y {
			return true
		}
	}
	return false
}

func flatten(stack [][]string) []string {
	s := make([]string, 0)
	for _, elem := range stack {
		var token string
		if len(elem) > 1 {
			token = elem[0] + "(" + strings.Join(elem[1:], " ") + ")"
		} else {
			token = elem[0]
		}
		s = append(s, token)
	}

	return s
}

// ./fetch http://www.w3.org/TR/2006/REC-xml11-20060816 | go run . div div1 enumar li p span
