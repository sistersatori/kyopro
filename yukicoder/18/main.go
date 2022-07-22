package main

import (
	"bufio"
	"fmt"
	"os"
)

// func solve(r *bufio.Reader, w *bufio.Writer) {}

func rotN(c, n byte) byte {
	var a, z byte
	switch {
	case 'a' <= c && c <= 'z':
		a, z = 'a', 'z'
	case 'A' <= c && c <= 'Z':
		a, z = 'A', 'Z'
	default:
		return c
	}
	return (c-a+n)%(z-a+1) + a
}

func reverseRotN(c, n byte) byte {
	var a, z byte
	switch {
	case 'a' <= c && c <= 'z':
		a, z = 'a', 'z'
	case 'A' <= c && c <= 'Z':
		a, z = 'A', 'Z'
	default:
		return c
	}
	fmt.Println(c, a, n, z, a)
	fmt.Println(c-a-n, z-a+1, a)
	fmt.Println((c - a - n) % (z - a + 1))
	return (c-a-n)%(z-a+1) + a
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var s string
	fmt.Fscan(r, &s)

	for i := 0; i < len(s); i++ {
		fmt.Fprint(w, string(reverseRotN(s[i], uint8(i+1))))
	}
	fmt.Fprintln(w)
}
