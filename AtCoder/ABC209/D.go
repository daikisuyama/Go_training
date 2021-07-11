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

	var n, q int
	fmt.Fscan(scin, &n, &q)
	var tree [][]int = make([][]int, n)
	for i := 0; i < n-1; i++ {
		var a, b int
		fmt.Fscan(scin, &a, &b)
		tree[a-1] = append(tree[a-1], b-1)
		tree[b-1] = append(tree[b-1], a-1)
	}
	var check []int = make([]int, n)
	check[0] = 1
	var now []int = make([]int, 0)
	now = append(now, 0)
	for len(now) > 0 {
		x := now[0]
		now = now[1:]
		for _, nex := range tree[x] {
			if check[nex] == 0 {
				check[nex] = -check[x]
				now = append(now, nex)
			}
		}
	}
	for i := 0; i < q; i++ {
		var c, d int
		fmt.Fscan(scin, &c, &d)
		if check[c-1]+check[d-1] == 0 {
			fmt.Fprintln(prout, "Road")
		} else {
			fmt.Fprintln(prout, "Town")
		}
	}
}
