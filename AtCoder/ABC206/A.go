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

	var n int
	fmt.Fscan(scin, &n)
	x := int(math.Floor(1.08 * float64(n)))

	switch {
	case x < 206:
		fmt.Fprintln(prout, "Yay!")
	case x > 206:
		fmt.Fprintln(prout, ":(")
	default:
		fmt.Fprintln(prout, "so-so")
	}
}
