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
	var h, w, x, y int
	fmt.Fscan(scin, &h, &w, &x, &y)
	s := make([]string, h)
	for i := 0; i < h; i++ {
		fmt.Fscan(scin, &s[i])
	}
	ans := 0
	for i := x - 1; i < h; i++ {
		if s[i][y-1] == '.' {
			ans++
		} else {
			break
		}
	}
	for i := x - 1; i >= 0; i-- {
		if s[i][y-1] == '.' {
			ans++
		} else {
			break
		}
	}
	for i := y - 1; i < w; i++ {
		if s[x-1][i] == '.' {
			ans++
		} else {
			break
		}
	}
	for i := y - 1; i >= 0; i-- {
		if s[x-1][i] == '.' {
			ans++
		} else {
			break
		}
	}
	fmt.Fprintln(prout, ans-3)
}
