package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	scin := bufio.NewReader(os.Stdin)
	prout := bufio.NewWriter(os.Stdout)
	defer prout.Flush()

	var x []int = make([]int, 3)
	fmt.Fscan(scin, &x[0], &x[1], &x[2])
	sort.Slice(x, func(i, j int) bool { return x[i] > x[j] })
	fmt.Fprintln(prout, x[0]+x[1])
}
