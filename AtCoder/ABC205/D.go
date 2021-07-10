package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	scin := bufio.NewReader(os.Stdin)
	prout := bufio.NewWriter(os.Stdout)
	defer prout.Flush()

	var n, q int
	fmt.Fscan(scin, &n, &q)
	var a_tmp map[int]int = make(map[int]int)
	var tmp int
	for i := 0; i < n; i++ {
		fmt.Fscan(scin, &tmp)
		a_tmp[tmp] = 1
	}
	var a []int64 = make([]int64, 0)
	a = append(a, 0)
	for key := range a_tmp {
		a = append(a, int64(key))
	}
	a = append(a, 2*int64(math.Pow10(18)))
	n = len(a)
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	var b []int64 = make([]int64, n-1)
	for i := 0; i < n-1; i++ {
		b[i] = a[i+1] - a[i] - 1
		if i > 0 {
			b[i] += b[i-1]
		}
	}
	for i := 0; i < q; i++ {
		var k int64
		fmt.Fscan(scin, &k)
		j := sort.Search(n-1, func(i int) bool { return k <= b[i] })
		if j == 0 {
			fmt.Fprintln(prout, k)
		} else {
			fmt.Fprintln(prout, a[j]+(k-b[j-1]))
		}
	}
}
