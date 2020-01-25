package main

import (
	"bytes"
	"fmt"
)

// IntSet 은 소수의 양의 정수 집합니다.
// 제로 값은 진 집합을 나타낸다.
type IntSet struct {
	words []uint64
}

// Has 는 집합에 양의 값 x가 있는지 여부를 보고한다
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add 는 양의 값 x를 집합에 추가한다
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
	fmt.Printf("%08b\n", s.words[word])
}

// UnionWith 는 s를 s와 t의 합집합으로 설정한다
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String 은 집합을 "{1 2 3}" 형태의 문자열로 반환한다
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String())

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String())

	x.UnionWith(&y)
	fmt.Println(x.String())

	fmt.Println(x.Has(9), x.Has(123))

	fmt.Println(&x)
	fmt.Println(x.String())
	fmt.Println(x)
}
