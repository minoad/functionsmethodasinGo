package main

import (
	"fmt"
	"testing"
)

var testSampleLarge []int = []int{1, 3, 2, 4, 8, 6, 7, 2, 3, 0}
var testSampleSmall []int = []int{3, 2, 1}

func TestBubbleSort(t *testing.T) {
	fmt.Println(testSampleLarge)
	BubbleSort(&testSampleLarge)
	fmt.Println(testSampleLarge)
}

func Test_collectInt(t *testing.T) {
	collectInt()
}

func Test_prepareList(t *testing.T) {
	prepareList("0 41 2 1 8 7")
}
