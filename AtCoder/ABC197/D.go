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
	var n, x0, y0, xn2, yn2 float64
	fmt.Fscan(scin, &n)
	fmt.Fscan(scin, &x0, &y0)
	fmt.Fscan(scin, &xn2, &yn2)
	p0 := complex(x0, y0)
	pn2 := complex(xn2, yn2)
	th := complex(math.Cos(math.Pi*2/n), math.Sin(math.Pi*2/n))
	q := (p0 + pn2) / 2
	ans := (p0-q)*th + q
	fmt.Fprintln(prout, real(ans), imag(ans))
}
