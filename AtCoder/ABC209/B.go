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

	var n, x int
	fmt.Fscan(scin, &n, &x)
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(scin, &a)
		if i%2 == 0 {
			x -= a
		} else {
			x -= a
			x++
		}
	}
	if x < 0 {
		fmt.Fprintln(prout, "No")
	} else {
		fmt.Fprintln(prout, "Yes")
	}
}
