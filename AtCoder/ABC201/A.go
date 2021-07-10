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

	var a []int = make([]int, 3)
	fmt.Fscan(scin, &a[0], &a[1], &a[2])
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	if a[2]-a[1] == a[1]-a[0] {
		fmt.Fprintln(prout, "Yes")
	} else {
		fmt.Fprintln(prout, "No")
	}
}
