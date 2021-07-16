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
	var a, b, w int
	fmt.Fscan(scin, &a, &b, &w)
	w *= 1000
	// みかんの個数をループ
	ans1, ans2 := 1<<61, 0
	for i := 0; i <= 1000000; i++ {
		if a*i <= w && w <= b*i {
			if i < ans1 {
				ans1 = i
			}
			if i > ans2 {
				ans2 = i
			}
		}
	}
	if ans2 == 0 {
		fmt.Fprintln(prout, "UNSATISFIABLE")
	} else {
		fmt.Fprintln(prout, ans1, ans2)
	}
}
