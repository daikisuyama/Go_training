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
	var x string
	fmt.Fscan(scin, &x)
	var y []string
	for _, s := range x {
		if s == '.' {
			break
		}
		y = append(y, string(s))
	}
	fmt.Fprintln(prout, strings.Join(y, ""))
}
