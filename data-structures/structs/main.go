package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {

	man := Person{Name: "queen", Age: 12}
	// output will be {queen 12}
	fmt.Printf("%v\n", man)
	//output will be {Name:queen Age:12}
	fmt.Printf("%+v\n", man)

}
