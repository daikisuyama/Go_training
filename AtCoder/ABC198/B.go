package main

import (
	"bufio"
	"fmt"
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
	for n%10 == 0 {
		n /= 10
		if n == 0 {
			break
		}
	}
	m := strconv.Itoa(n)
	l := len(m)
	flag := false
	for i := 0; i < l/2; i++ {
		if m[i] != m[l-1-i] {
			flag = true
		}
	}
	if flag {
		fmt.Fprintln(prout, "No")
	} else {
		fmt.Fprintln(prout, "Yes")
	}
}
