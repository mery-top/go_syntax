package main

import "fmt"

func main(){
	x:=10

	var p *int = &x
	fmt.Println(*p)
}