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

	var n int
	fmt.Fscan(scin, &n)
	var now int
	var ans int = 1
	for ; now < n; ans++ {
		now += ans
	}
	fmt.Fprintln(prout, ans-1)
}
