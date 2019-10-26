package main

import "fmt"

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
}
