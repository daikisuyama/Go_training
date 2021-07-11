package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scin := bufio.NewReader(os.Stdin)
	prout := bufio.NewWriter(os.Stdout)
	defer prout.Flush()

	// code
	// 前半後半に最初から分ける
	var first, second []byte
	var n int
	var s string
	fmt.Fscan(scin, &n)
	fmt.Fscan(scin, &s)
	for i := 0; i < 2*n; i++ {
		if i < n {
			first = append(first, s[i])
		} else {
			second = append(second, s[i])
		}
	}
	var q int
	fmt.Fscan(scin, &q)
	for i := 0; i < q; i++ {
		var t, a, b int
		fmt.Fscan(scin, &t, &a, &b)
		a--
		b--

		if t == 2 {
			first, second = second, first
		} else {
			switch {
			case a < n && b < n:
				first[a], first[b] = first[b], first[a]
			case a < n:
				first[a], second[b-n] = second[b-n], first[a]
			case b < n:
				second[a-n], first[b] = first[b], second[a-n]
			default:
				second[a-n], second[b-n] = second[b-n], second[a-n]
			}
		}
	}
	var ans []string
	for i := 0; i < n; i++ {
		ans = append(ans, string(first[i]))
	}
	for i := 0; i < n; i++ {
		ans = append(ans, string(second[i]))
	}
	fmt.Fprintln(prout, strings.Join(ans, ""))
}
