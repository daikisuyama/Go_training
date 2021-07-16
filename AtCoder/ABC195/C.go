package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	scin := bufio.NewReader(os.Stdin)
	prout := bufio.NewWriter(os.Stdout)
	defer prout.Flush()

	// code
	// それぞれの桁でいくつあるか
	// i桁の場合は[(i-1)/3]
	// i桁は100…00から999…99
	var n int
	fmt.Fscan(scin, &n)
	a := make([]int, 17)
	a[0] = 1
	for i := 0; i < 16; i++ {
		a[i+1] = a[i] * 10
	}
	ans := 0
	for i := 1; i < 17; i++ {
		if n <= a[i]-1 {
			ans += int(math.Floor(float64(i-1)/3)) * (n - a[i-1] + 1)
			break
		}
		ans += int(math.Floor(float64(i-1)/3)) * (a[i] - 1 - a[i-1] + 1)
	}
	fmt.Fprintln(prout, ans)
}
