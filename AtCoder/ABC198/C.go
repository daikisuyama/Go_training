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

	// code
	var r, x, y float64
	fmt.Fscan(scin, &r, &x, &y)
	z := int(math.Ceil(math.Sqrt(x*x+y*y) / r))
	if z == 1 && x*x+y*y != r*r {
		fmt.Fprintln(prout, 2)
	} else {
		fmt.Fprintln(prout, z)
	}
}
