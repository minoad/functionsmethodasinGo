package main

import (
	"fmt"
	"testing"
)

// For example, let’s say that I want to assume the following values for acceleration, initial velocity, and initial displacement: a = 10, vo = 2, so = 1. I can use the following statement to call GenDisplaceFn() to generate a function fn which will compute displacement as a function of time.
// fn := GenDisplaceFn(10, 2, 1)
// Then I can use the following statement to print the displacement after 3 seconds.
// fmt.Println(fn(3))
// And I can use the following statement to print the displacement after 5 seconds.
// fmt.Println(fn(5))

func TestGenDisplaceFn(t *testing.T) {
	n := GenDisplaceFn(10, 2, 10)
	x := n(3)
	x1 := n(5)
	x2 := n(50)

	fmt.Printf("%v -- %v -- %v --", x, x1, x2)
}

func Test_collectDisplacementData(t *testing.T) {
	collectDisplacementData()
}
