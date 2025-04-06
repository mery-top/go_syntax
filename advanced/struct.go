package main

import "fmt"

type Person struct{
	name string
	age int
}

func main(){
	p:=Person{name: "Hogn", age:13}
	fmt.Println(p.name, p.age)
}