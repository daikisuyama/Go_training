package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scin := bufio.NewReader(os.Stdin)
	prout := bufio.NewWriter(os.Stdout)
	defer prout.Flush()

	var a, b, c int
	fmt.Fscan(scin, &a, &b, &c)
	switch {
	case a == b:
		fmt.Fprintln(prout, c)
	case b == c:
		fmt.Fprintln(prout, a)
	case a == c:
		fmt.Fprintln(prout, b)
	default:
		fmt.Fprintln(prout, 0)
	}
}
