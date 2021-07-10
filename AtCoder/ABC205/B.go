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

	var n int
	fmt.Fscan(scin, &n)
	var a []int = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(scin, &a[i])
	}
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	var flag bool
	for i := 0; i < n; i++ {
		if a[i] != i+1 {
			flag = true
			break
		}
	}
	if flag {
		fmt.Fprintln(prout, "No")
	} else {
		fmt.Fprintln(prout, "Yes")
	}
}
