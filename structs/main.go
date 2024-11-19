package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (n *Person) UpdateName(name string) {
	n.Name = name
}

func main() {

	name := Person{
		Name: "Quinter",
		Age:  12,
	}
	//fmt.Printf("%v") : output of struct will be {"Quinter 12"}
	//using %+v gives  this output includes fields of structs and assigned values fmt.printf("%+v") {Name:Quinter Age:12}
	fmt.Printf("first instance of struct %+v\n", name)
	name.UpdateName("Apondi")
	fmt.Printf("second intance of struct after struct %+v\n", name)

}
