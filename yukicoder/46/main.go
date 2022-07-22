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

	var a, b int
	fmt.Fscan(r, &a, &b)

	if b%a == 0 {
		fmt.Fprintln(w, b/a)
	} else {
		fmt.Fprintln(w, b/a+1)
	}
}
