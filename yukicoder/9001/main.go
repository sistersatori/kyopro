package main

import (
	"bufio"
	"fmt"
	"os"
)

// func solve(r *bufio.Reader, w *bufio.Writer) {}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var A, B int
	var S string

	fmt.Fscan(r, &A, &B, &S)
	fmt.Fprintln(w, A+B, S)
}
