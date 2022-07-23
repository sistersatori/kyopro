package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// func solve(r *bufio.Reader, w *bufio.Writer) {}

func main() {
	// r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	loc, _ := time.LoadLocation("Asia/Tokyo")
	d := time.Date(2015, time.January, 1, 0, 0, 0, 0, loc)

	var count int
	for i := 0; i < 365; i++ {
		month := int(d.Month())
		day := int(d.Day())
		daywa := day%10 + day/10

		if month == daywa {
			count++
		}
		d = d.AddDate(0, 0, 1)
	}
	fmt.Fprintln(w, count)
}
