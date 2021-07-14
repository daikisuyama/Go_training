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
	var a, b, c, d int
	fmt.Fscan(scin, &a, &b)
	fmt.Fscan(scin, &c, &d)
	fmt.Fprintln(prout, b-c)
}
