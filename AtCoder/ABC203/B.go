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

	var n, k int
	fmt.Fscan(scin, &n, &k)
	var ans int
	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			ans += 100 * i
			ans += j
		}
	}
	fmt.Fprintln(prout, ans)
}
