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

	var N int
	fmt.Fscan(r, &N)

	var A, sum uint
	for i := 0; i < N; i++ {
		fmt.Fscan(r, &A)
		sum += A
	}

	fmt.Fprintln(w, sum)
}
