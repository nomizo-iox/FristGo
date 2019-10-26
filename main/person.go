package main

import ("fmt"
)

// Go does not have classes, so we use Struct
type person struct {
	first_name string
	last_name string
	email string
	age int
}

func main()  {
	p := person{first_name: "Samuel", last_name: "Ademola", email: "sa10@gmail.com", age:32}
	fmt.Println(p)
	fmt.Println(p.last_name)
}