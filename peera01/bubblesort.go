package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("a ...interface{}")
	data := collectInt()
	BubbleSort(&data)
	fmt.Println(data)
}

func collectInt() []int {
	// Collect the data
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter a list of up to 10 integers seperated by spaces.")
	fmt.Println("-----------------------------------------")

	fmt.Print("-->")
	text, _ := reader.ReadString('\n')

	// prepare the data
	n, err := prepareList(strings.Replace(text, "\n", "", -1))
	if err != nil {
		log.Fatalf("unable to prepare list.  Bailing out. %v", err)
	}
	return n
}

func prepareList(intList string) ([]int, error) {
	// Split it to get our slice
	strSlice := strings.Split(intList, " ")

	// create a []int sized correctly
	intSlice := make([]int, len(strSlice))

	// loop through the split strings and convert them to ints and add to the int slice
	// TODO: right now i simply bail.  I should look for corner cases like blank space or tab.
	for i, val := range strSlice {
		n, err := strconv.Atoi(val)
		if err != nil {
			return nil, fmt.Errorf("unable to convert %s to integer: %v", val, err)
		}
		intSlice[i] = n
	}
	return intSlice, nil
}

// BubbleSort accepts a slice of integers, sorts it and then returns.
func BubbleSort(n *[]int) {
	// Up to len(n) number of times, loop over the list.
	for it := 0; it <= len(*n); it++ {

		for i := 0; i < len(*n)-1; i++ {
			// compare the nth element to the nth+1 element.
			if (*n)[i] > (*n)[i+1] {
				// if the nth+1 element is smaller than the nth element, swap.
				Swap(n, i)
			}
		}
	}
}

// Swap Simply swap the values
func Swap(n *[]int, index int) {
	(*n)[index], (*n)[index+1] = (*n)[index+1], (*n)[index]
}
