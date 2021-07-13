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
	var s string
	fmt.Fscan(scin, &s)
	fmt.Fprintln(prout, s[1:]+string(s[0]))
	//fmt.Fprintln(prout, s[1:]+s[:1])
}
