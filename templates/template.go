package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var h, w string
	fmt.Fscan(in, &h, &w)
	fmt.Fprintln(out, h, w)
}
