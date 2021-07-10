package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type kokumin struct {
	ind, bangou int
}

func main() {
	scin := bufio.NewReader(os.Stdin)
	prout := bufio.NewWriter(os.Stdout)
	defer prout.Flush()
	var n, k int
	fmt.Scan(&n, &k)
	var a []kokumin = make([]kokumin, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(scin, &(a[i].bangou))
		a[i].ind = i
	}
	sort.Slice(a, func(i, j int) bool { return a[i].bangou < a[j].bangou })
	var ans []int = make([]int, n)
	for i := 0; i < n; i++ {
		if i < k%n {
			ans[a[i].ind] = k/n + 1
		} else {
			ans[a[i].ind] = k / n
		}
	}
	for i := 0; i < n; i++ {
		fmt.Fprintln(prout, ans[i])
	}
}
