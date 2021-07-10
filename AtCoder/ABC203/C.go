package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type per struct {
	a, b int64
}

func main() {
	scin := bufio.NewReader(os.Stdin)
	prout := bufio.NewWriter(os.Stdout)
	defer prout.Flush()

	var n, k int64
	fmt.Fscan(scin, &n, &k)
	var person map[int64]int64 = make(map[int64]int64)
	for i := 0; i < int(n); i++ {
		var a, b int64
		fmt.Fscan(scin, &a, &b)
		person[a] += b
	}
	var x []per = make([]per, 0)
	for a, b := range person {
		x = append(x, per{a, b})
	}
	sort.Slice(x, func(i, j int) bool { return x[i].a < x[j].a })
	n = int64(len(x))
	for i := 0; i < int(n); i++ {
		if k >= x[i].a {
			k += x[i].b
		} else {
			break
		}
	}
	fmt.Fprintln(prout, int(math.Min(math.Pow10(100), float64(k))))
}
