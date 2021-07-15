package main

import "fmt"

func main() {
	s := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}}

	// t1はshallow copy
	t1 := make([][]int, 2)
	for i := range s {
		t1[i] = make([]int, 4)
	}
	copy(t1, s)

	// t2はdeep copy
	t2 := make([][]int, 2)
	for i := range s {
		t2[i] = make([]int, 4)
		copy(t2[i], s[i])
	}

	// 要素の変更
	t1[0][0] = 100
	t2[0][0] = 99

	//出力：[[100 2 3 4] [5 6 7 8]]
	fmt.Println(s)

	//出力：[[100 2 3 4] [5 6 7 8]]
	fmt.Println(t1)

	//出力：[[99 2 3 4] [5 6 7 8]]
	fmt.Println(t2)
}
