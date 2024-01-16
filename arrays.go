package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp: ", a) // emp [0 0 0 0 0]

	a[4] = 100
	fmt.Println("Set: ", a)    // [0 0 0 0 100]
	fmt.Println("Get: ", a[4]) // 100

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl: ", b) // [1 2 3 4 5]

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD) //2d: [[0 1 2][1 2 3]]
}
