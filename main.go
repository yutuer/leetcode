package main

import "fmt"

const(
	i = 3
	a byte = 100
	b int = 1e20  // constant 100000000000000000000 overflows int
)

func main() {
	fmt.Println("test")
	fmt.Println(i, a, b)
}
