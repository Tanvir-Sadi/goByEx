package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s) // constant

	const n = 500000

	const d = 3e20 / n
	fmt.Println(d)        // 6e+14
	fmt.Println(int64(d)) // 600000000000000

	fmt.Println(math.Sin(n)) // 0.17783120151825887
}
