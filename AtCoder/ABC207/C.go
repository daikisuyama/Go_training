package main

import (
	"bufio"
	"fmt"
	"os"
)

type kuk struct {
	l, r int
}

func main() {
	scin := bufio.NewReader(os.Stdin)
	prout := bufio.NewWriter(os.Stdout)
	defer prout.Flush()

	var n int
	fmt.Fscan(scin, &n)
	var tlr []kuk = make([]kuk, n)
	for i := 0; i < n; i++ {
		var t int
		fmt.Fscan(scin, &t, &(tlr[i].l), &(tlr[i].r))
		tlr[i].l *= 3
		tlr[i].r *= 3
		if t == 4 || t == 3 {
			tlr[i].l += 1
		}
		if t == 4 || t == 2 {
			tlr[i].r -= 1
		}
	}
	var ans int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			var check bool
			if tlr[i].l < tlr[j].l {
				if tlr[i].r >= tlr[j].l {
					check = true
				}
			} else if tlr[j].l < tlr[i].l {
				if tlr[j].r >= tlr[i].l {
					check = true
				}
			} else {
				check = true
			}
			if check {
				ans++
			}
		}
	}
	fmt.Fprintln(prout, ans)
}
