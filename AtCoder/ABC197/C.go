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

	// code
	var n int
	fmt.Fscan(scin, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(scin, &a[i])
	}
	ans := 1 << 62
	for i := 0; i < (1 << n); i++ {
		flag := 0
		ans_s := 0
		ans_s_s := 0
		for j := 0; j < n; j++ {
			if flag == (i>>j)&1 {
				ans_s_s |= a[j]
			} else {
				ans_s ^= ans_s_s
				ans_s_s = a[j]
			}
			if j == n-1 {
				ans_s ^= ans_s_s
			}
		}
		if ans_s < ans {
			ans = ans_s
		}
	}
	fmt.Fprintln(prout, ans)
}
