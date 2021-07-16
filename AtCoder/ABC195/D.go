package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type wv struct {
	w, v int
}

func main() {
	scin := bufio.NewReader(os.Stdin)
	prout := bufio.NewWriter(os.Stdout)
	defer prout.Flush()

	// code
	var n, m, q int
	fmt.Fscan(scin, &n, &m, &q)
	bag := make([]wv, n)
	x := make([]int, m)
	for i := 0; i < n; i++ {
		fmt.Fscan(scin, &bag[i].w, &bag[i].v)
	}
	sort.Slice(bag, func(i, j int) bool { return bag[i].w < bag[j].w })
	for i := 0; i < m; i++ {
		fmt.Fscan(scin, &x[i])
	}
	for i := 0; i < q; i++ {
		var l, r int
		fmt.Fscan(scin, &l, &r)
		// 選べる箱
		var y []int
		for i := 0; i < m; i++ {
			if i < l-1 || r-1 < i {
				y = append(y, x[i])
			}
		}
		sort.Slice(y, func(i, j int) bool { return y[i] < y[j] })
		// 他に余ってる荷物
		z := make([]wv, n)
		copy(z, bag)
		// 入れれる荷物の候補
		cand := make([]wv, 0)
		ans := 0
		// 選べる箱
		for _, j := range y {
			for len(z) > 0 {
				if z[0].w <= j {
					cand = append(cand, z[0])
					z = z[1:]
				} else {
					break
				}
			}
			sort.Slice(cand, func(i, j int) bool { return cand[i].v > cand[j].v })
			if len(cand) == 0 {
				continue
			} else {
				ans += cand[0].v
				cand = cand[1:]
			}
		}
		fmt.Fprintln(prout, ans)
	}
}
