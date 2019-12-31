package main

import (
	"fmt"
	"testing"
)

func Test_main(t *testing.T) {
	main()
}

func x1(x int) int {
	return x + 1
}
func Test_applyIt(t *testing.T) {
	// anonymous function
	fmt.Println(applyIt(func(x int) int { return x + 1 }, 7))
	// defined function
	fmt.Println(applyIt(x1, 2))
	// variable as a function
	n := x1
	fmt.Println(applyIt(n, 45))
}

func Test_makeDistOrigin(t *testing.T) {
	n := makeDistOrigin(4.5, 7.0)
	fmt.Println(n(2.6, 45.4))
}

func Test_getMax(t *testing.T) {
	fmt.Println(getMax(1, 4, 3, 6, 78, 2, 8))
}

// func Test_fA(t *testing.T) {
// 	fA(func(x int) string { return "a ...interface{}" })
// }

func Test_main2(t *testing.T) {
	main2()
}
