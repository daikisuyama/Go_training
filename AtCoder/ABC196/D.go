package main

import (
	"bufio"
	"fmt"
	"os"
)

var h, w, a, b int

// 新しいgrid作らないと参照してしまうので注意
// 与えられたグリッドの(i,j)を決めて、次を決める関数
// t1:2*1の畳の枚数
func search(grid [][]int, i int, j int, t1 int) int {
	if i == h-1 && j == w-1 {
		// 最終マスの時
		if t1 == a {
			return 1
		} else {
			return 0
		}
	} else if i == h-1 {
		// 最終行の時（縦方向はない）
		if grid[i][j] == 1 || grid[i][j+1] == 1 {
			grid_n := make([][]int, h)
			for i := range grid {
				grid_n[i] = make([]int, w)
				copy(grid_n[i], grid[i])
			}
			return search(grid_n, i, j+1, t1)
		} else {
			grid_n1 := make([][]int, h)
			grid_n2 := make([][]int, h)
			for i := range grid {
				grid_n1[i] = make([]int, w)
				grid_n2[i] = make([]int, w)
				copy(grid_n1[i], grid[i])
				copy(grid_n2[i], grid[i])
			}
			// 横で2枚
			grid_n1[i][j] = 1
			grid_n1[i][j+1] = 1
			return search(grid_n1, i, j+1, t1+1) + search(grid_n2, i, j+1, t1)
		}
	} else if j == w-1 {
		// 最終列の時（横方向はない）
		if grid[i][j] == 1 {
			grid_n := make([][]int, h)
			for i := range grid {
				grid_n[i] = make([]int, w)
				copy(grid_n[i], grid[i])
			}
			return search(grid_n, i+1, 0, t1)
		} else {
			grid_n1 := make([][]int, h)
			grid_n2 := make([][]int, h)
			for i := range grid {
				grid_n1[i] = make([]int, w)
				grid_n2[i] = make([]int, w)
				copy(grid_n1[i], grid[i])
				copy(grid_n2[i], grid[i])
			}
			// 縦で2枚
			grid_n1[i][j] = 1
			grid_n1[i+1][j] = 1
			return search(grid_n1, i+1, 0, t1+1) + search(grid_n2, i+1, 0, t1)
		}
	} else {
		// その他（縦方向も横方向もあり得る）
		if grid[i][j] == 1 {
			// そのマスがダメ
			grid_n := make([][]int, cap(grid))
			for i := range grid {
				grid_n[i] = make([]int, w)
				copy(grid_n[i], grid[i])
			}
			return search(grid_n, i, j+1, t1)
		} else if grid[i][j+1] == 1 {
			// 横方向がダメ
			grid_n1 := make([][]int, h)
			grid_n2 := make([][]int, h)
			for i := range grid {
				grid_n1[i] = make([]int, w)
				grid_n2[i] = make([]int, w)
				copy(grid_n1[i], grid[i])
				copy(grid_n2[i], grid[i])
			}
			// 縦で2枚
			grid_n1[i][j] = 1
			grid_n1[i+1][j] = 1
			return search(grid_n1, i, j+1, t1+1) + search(grid_n2, i, j+1, t1)
		} else {
			// どちらもOK
			grid_n1 := make([][]int, h)
			grid_n2 := make([][]int, h)
			grid_n3 := make([][]int, h)
			for i := range grid {
				grid_n1[i] = make([]int, w)
				grid_n2[i] = make([]int, w)
				grid_n3[i] = make([]int, w)
				copy(grid_n1[i], grid[i])
				copy(grid_n2[i], grid[i])
				copy(grid_n3[i], grid[i])
			}
			// 横で2枚
			grid_n1[i][j] = 1
			grid_n1[i][j+1] = 1
			// 縦で2枚
			grid_n2[i][j] = 1
			grid_n2[i+1][j] = 1
			return search(grid_n1, i, j+1, t1+1) + search(grid_n2, i, j+1, t1+1) + search(grid_n3, i, j+1, t1)
		}
	}
}

func main() {
	scin := bufio.NewReader(os.Stdin)
	prout := bufio.NewWriter(os.Stdout)
	defer prout.Flush()

	// code
	fmt.Fscan(scin, &h, &w, &a, &b)
	grid := make([][]int, h)
	for i := 0; i < h; i++ {
		grid[i] = make([]int, w)
	}
	fmt.Fprintln(prout, search(grid, 0, 0, 0))
}
