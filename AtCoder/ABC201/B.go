package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type yama struct {
	s string
	t int
}

func main() {
	scin := bufio.NewReader(os.Stdin)
	prout := bufio.NewWriter(os.Stdout)
	defer prout.Flush()

	var n int
	fmt.Fscan(scin, &n)
	var yamas []yama = make([]yama, n)
	for i := 0; i < n; i++ {
		var s string
		var t int
		fmt.Fscan(scin, &s, &t)
		yamas = append(yamas, yama{s, t})
	}
	sort.Slice(yamas, func(i, j int) bool { return yamas[i].t > yamas[j].t })
	fmt.Fprintln(prout, yamas[1].s)
}
