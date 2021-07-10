package main

import "fmt"

func main() {
	var p int
	fmt.Scan(&p)
	coin := make([]int, 10)
	coin[0] = 1
	for i := 2; i <= 10; i++ {
		coin[i-1] = coin[i-2] * i
	}
	co := 0
	for i := 9; i > 0; i-- {
		co += p / coin[i]
		p %= coin[i]
	}
	fmt.Println(co + p)
}
