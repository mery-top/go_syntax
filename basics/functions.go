package main

import "fmt"

func greet(name string) string{
	return "HI"+name
}

func sum(v1 int, v2 int)int{
	return v1+v2
}

func main(){
	fmt.Println(greet("Alcie"))
	fmt.Println(sum(2,3))
}