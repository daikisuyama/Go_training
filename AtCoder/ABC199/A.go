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

	// code
	var a, b, c int
	fmt.Fscan(scin, &a, &b, &c)
	if a*a+b*b < c*c {
		fmt.Fprintln(prout, "Yes")
	} else {
		fmt.Fprintln(prout, "No")
	}
}
