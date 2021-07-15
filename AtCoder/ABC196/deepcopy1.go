package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4}

	// deepcopy
	t := make([]int, 4)
	copy(t, s)

	// 要素の変更
	t[1] = 100

	//出力：[1 2 3 4]
	fmt.Println(s)

	//出力：[1 100 3 4]
	fmt.Println(t)
}
