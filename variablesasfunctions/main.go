package main

import (
	"fmt"
	"math"
)

var funcVar func(int) int

func incFn(x int) int {
	return x + 1
}

func applyIt(afunct func(int) int, val int) int {
	return afunct(val)
}

// variadic
// can also pass a slice getMax([]int{1,2,3,4,5})
func getMax(vals ...int) int {
	maxV := -1
	for _, v := range vals {
		if v > maxV {
			maxV = v
		}
	}
	return maxV
}

func makeDistOrigin(o_x, o_y float64) func(float64, float64) float64 {
	// composes a function filled out with the origin detail.
	// returned function will need to take a new point for the destination.
	return func(x, y float64) float64 {
		return math.Sqrt(math.Pow(x-o_x, 2) + math.Pow(y-o_y, 2))
	}
}

var f func(string) int

func test(x string) int {
	return len(x)
}

// func fA(Fb func(int) string) {
// 	fmt.Println("a")
// }

func fA() func() int {
	i := 0
	return func() int {
			i++
			return i
	}
}
func main2() {
 fB := fA()
 fmt.Print(fB())
 fmt.Print(fB())
}

func main() {
	// no parens after incFn because we are assigning, not calling a function.
	funcVar = incFn
	fmt.Println(funcVar(1))
	f = test
	fmt.Println(f("hello"))
}
