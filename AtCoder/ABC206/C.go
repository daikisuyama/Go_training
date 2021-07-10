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

	var n int
	fmt.Fscan(scin, &n)
	var a map[int]int = make(map[int]int)
	for i := 0; i < n; i++ {
		var j int
		fmt.Fscan(scin, &j)
		a[j]++
	}
	var ans int
	for _, v := range a {
		ans += v * (v - 1) / 2
	}
	fmt.Fprintln(prout, n*(n-1)/2-ans)
}
