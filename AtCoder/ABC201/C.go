package main

import (
	"bufio"
	"fmt"
	"os"
)

// nPk
// 「 k<0 or n>k 」の場合は0として定義
func permutation(n int, k int) (ret int) {
	if 0 <= k && k <= n {
		ret = 1
		for i := 0; i < k; i++ {
			ret *= n - i
		}
	}
	return
}

// n!
func factorial(n int) int {
	return permutation(n, n)
}

// nCk
func combination(n int, k int) int {
	return permutation(n, k) / factorial(k)
}

func main() {
	scin := bufio.NewReader(os.Stdin)
	prout := bufio.NewWriter(os.Stdout)
	defer prout.Flush()

	var x, y, z int
	var s string
	fmt.Fscan(scin, &s)
	for i := 0; i < 10; i++ {
		if s[i] == 'o' {
			x++
		} else if s[i] == 'x' {
			z++
		} else {
			y++
		}
	}
	var ans int
	if x <= 1 && x+y >= 1 {
		ans += combination(y, 1-x) * 1
	}
	if x <= 2 && x+y >= 2 {
		ans += combination(y, 2-x) * (24/6*2 + 24/2/2)
	}
	if x <= 3 && x+y >= 3 {
		ans += combination(y, 3-x) * (3 * 24 / 2)
	}
	if x <= 4 && x+y >= 4 {
		ans += combination(y, 4-x) * (24)
	}
	fmt.Fprintln(prout, ans)

}
