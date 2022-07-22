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

	for i := 1; i <= N; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Fprintln(w, "FizzBuzz")
		case i%3 == 0:
			fmt.Fprintln(w, "Fizz")
		case i%5 == 0:
			fmt.Fprintln(w, "Buzz")
		default:
			fmt.Fprintln(w, i)
		}
	}
}
