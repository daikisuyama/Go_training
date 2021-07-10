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

	var a, b, c, d int
	fmt.Fscan(scin, &a, &b, &c, &d)
	if b >= d*c {
		fmt.Fprintln(prout, -1)
	} else {
		fmt.Fprintln(prout, math.Ceil(float64(a)/float64(d*c-b)))
	}
}
