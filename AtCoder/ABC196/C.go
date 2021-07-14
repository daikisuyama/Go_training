package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	scin := bufio.NewReader(os.Stdin)
	prout := bufio.NewWriter(os.Stdout)
	defer prout.Flush()

	// code
	var n int
	fmt.Fscan(scin, &n)
	l := len(strconv.Itoa(n))
	ans := 0
	if l%2 == 1 {
		for m := 1; m <= l/2; m++ {
			ans += (9 * int(math.Pow10(m-1)))
		}
	} else {
		for m := 1; m < l/2; m++ {
			ans += (9 * int(math.Pow10(m-1)))
		}
		for x := int(math.Pow10(l/2 - 1)); x*int(math.Pow10(l/2))+x <= n; x++ {
			ans++
		}
	}
	fmt.Fprintln(prout, ans)
}
