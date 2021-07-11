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

	var a, b int
	fmt.Fscan(scin, &a, &b)
	if b < a {
		fmt.Fprintln(prout, 0)
	} else {
		fmt.Fprintln(prout, b-a+1)
	}
}
