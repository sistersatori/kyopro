package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var L, N int
	fmt.Fscan(in, &L, &N)

	W := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &W[i])
	}

	sort.SliceStable(W, func(i, j int) bool { return W[i] < W[j] })

	c, sum := 0, 0
	for i := 0; i < N; i++ {
		sum += W[i]
		if sum > L {
			break
		} else {
			c++
		}
	}

	fmt.Fprintln(out, c)
}
