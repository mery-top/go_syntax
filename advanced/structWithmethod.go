package main

import "fmt"

type Rectangle struct{
	width, height float64
}

func(r Rectangle) area() float64{
	return r.width * r.height
}

func main(){
	rect:= Rectangle{10, 8}
	fmt.Println(rect.area())
}
