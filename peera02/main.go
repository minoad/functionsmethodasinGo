package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// I created named types of float 64 in order to keep the variables straight.
type acceleration float64
type velocity float64
type displacement float64
type timeT float64

// GenDisplaceFn returns a closure over the displacement calculation
// These variables ARE float64's.  They are named types in order to keep them straight.
func GenDisplaceFn(a acceleration, vo velocity, so displacement) func(t timeT) float64 {
	// s =½ a t2 + vot + so

	// Likewise the t timeT is a float64.  Renamed in order to keep it straight.
	return func(t timeT) float64 {
		fmt.Printf("caluculating: .5 * %v + %v + %v + %v\n", float64(a), float64(t), float64(vo), float64(so))
		s := .5*float64(a)*float64(t) + float64(vo) + float64(so)
		return s
	}
}

// Helper function to get our floats based on the requested string.
func collectFloatValue(s string) float64 {
	// Collect the data
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please %s value. --> ", s)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("could not collect text %s: %v\n", text, err)
	}

	n, err := strconv.ParseFloat(strings.Replace(text, "\n", "", -1), 64)
	if err != nil {
		fmt.Printf("could not parse float %v: %v\n", n, err)
	}

	return n
}

// Collects the variables required for the closure.
func collectDisplacementData() (acceleration, velocity, displacement) {
	a := collectFloatValue("accelleration")
	vo := collectFloatValue("velocity")
	so := collectFloatValue("displacement")
	return acceleration(a), velocity(vo), displacement(so)
}

func main() {
	a, vo, so := collectDisplacementData()
	n := GenDisplaceFn(a, vo, so)
	for {
		t := collectFloatValue("time")
		fmt.Println(n(timeT(t)))
	}
}
