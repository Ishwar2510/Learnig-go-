package main

import (
	"fmt"
)

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev := Reverse(input)
	doubleRev := Reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q\n", rev)
	fmt.Printf("reversed again: %q\n", doubleRev)
}
func Reverse(s string) string {
	b := []byte(s)

	for start, end := 0, len(b)-1; start < end; start, end = start+1, end-1 {
		b[start], b[end] = b[end], b[start]
	}
	return string(b)
}
