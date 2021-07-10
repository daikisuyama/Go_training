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

	var a, b float64
	fmt.Fscan(scin, &a, &b)
	fmt.Fprintln(prout, a*b/100)
}
