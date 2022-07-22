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

	var S, F uint
	fmt.Fscan(r, &S, &F)
	fmt.Println(S/F + 1)
}
