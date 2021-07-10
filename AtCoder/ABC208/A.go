package main

import "fmt"

func main() {
	var a, b int16
	fmt.Scan(&a, &b)
	if a <= b && b <= 6*a {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
