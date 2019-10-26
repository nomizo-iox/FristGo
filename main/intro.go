package main

import (
	"errors"
	"fmt"
	"math"
)


func main() {

	// Assigning a String variable
	var firsName = "Sam"
	lastName := "Ademola"
	fmt.Println(firsName)
	fmt.Println(lastName)

	// Assigning an Int variable
	var x = 7
	y := 35
	fmt.Println(x)
	fmt.Println(y)

	// If Statement
	a := 32
	if a > 6 {
		fmt.Println("A is greater than 6")
	}

	// If Else Statement
	b := 24
	if b > 200 {
		fmt.Println("b is greater")
	} else if b < 200 {
		fmt.Println("b is smaller")
	} else {
		fmt.Println("The same")
	}

	// Static Array & Assignment
	var c [5]int
	fmt.Println(c)

	c[1] = 51
	fmt.Println(c)

	//Static Full declaration & initialization
	d := [5]int{12, 53, 78, 34, 24}
	fmt.Println(d)

	// Slicers (Dynamic Arrays)
	e := []string{"Samuel", "Abraham", "David"}
	fmt.Println(e)
	// This creates a new array in the background
	e = append(e, "Emmanuel")
	fmt.Println(e)

	// Maps for Go is like Dictionaries for Python, or HashMap for Java.  First it is k,v
	studentID := make(map[string]int)
	studentID["a001"] = 1
	studentID["a002"] = 2
	fmt.Println(studentID)
	fmt.Println("a001")
	fmt.Println(studentID["a002"])

	// Loops.  Go only has For Loop
	for i := 0; i < 5 ; i++ {
		fmt.Println(i)
	}

	// Loop over Array
	arr := []int{41, 9, 63, 10}
	for index, value := range arr {
		fmt.Println(index, value)
	}

	// Loop over Map
	mcu := make(map[int]string)
	mcu[1] = "Iron Man"
	mcu[2] = "Incredible Hulk"
	mcu[25] = "Black Panther"
	for key, value := range mcu {
		fmt.Println(key, value)
	}

	// Sum function
	result := sum(34, 24)
	fmt.Println(result)

	// Sqrt function
	result2, err := sqrt(16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result2)
	}
}

// Function with return value
func sum(o int, j int) int  {
	return o * j
}

// Function with multiple return value
func sqrt (k float64) (float64, error) {
	if k < 0 {
		return 0, errors.New("Undefine for negative numbers")
	}

	return math.Sqrt(k), nil
}

