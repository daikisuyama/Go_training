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
	var n int
	fmt.Fscan(scin, &n)
	a := 0
	b := 1001
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(scin, &x)
		if x > a {
			a = x
		}
	}
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(scin, &x)
		if x < b {
			b = x
		}
	}
	if a <= b {
		fmt.Fprintln(prout, b-a+1)
	} else {
		fmt.Fprintln(prout, 0)
	}
}
