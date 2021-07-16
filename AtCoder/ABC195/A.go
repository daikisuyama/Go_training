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
	var m, h int
	fmt.Fscan(scin, &m, &h)
	if h%m == 0 {
		fmt.Fprintln(prout, "Yes")
	} else {
		fmt.Fprintln(prout, "No")
	}
}
