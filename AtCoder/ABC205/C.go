package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	scin := bufio.NewReader(os.Stdin)
	prout := bufio.NewWriter(os.Stdout)
	defer prout.Flush()

	var a, b, c int
	fmt.Fscan(scin, &a, &b, &c)
	if c%2 == 0 {
		a = int(math.Abs(float64(a)))
		b = int(math.Abs(float64(b)))
	}
	switch {
	case a < b:
		fmt.Fprintln(prout, "<")
	case a > b:
		fmt.Fprintln(prout, ">")
	default:
		fmt.Fprintln(prout, "=")
	}
}
