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

	var n int
	fmt.Fscan(scin, &n)
	mod := int(math.Pow10(9)) + 7
	var c []int = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(scin, &c[i])
	}
	sort.Slice(c, func(i, j int) bool { return c[i] < c[j] })
	ans := 1
	flag := false
	for i := 0; i < n; i++ {
		if c[i] < i+1 {
			flag = true
			break
		} else {
			ans *= c[i] - i
			ans %= mod
		}
	}
	if flag {
		fmt.Fprintln(prout, 0)
	} else {
		fmt.Fprintln(prout, ans)
	}
}
