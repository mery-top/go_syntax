package main

import "fmt"

type Shape interface{
	area() float64
}

type Circle struct{
	width, height float64
}

func (c Circle) area() float64{
	return c.width * c.height
}

func interact(s Shape){
	fmt.Println(s.area())
}
func main(){
	c:= Circle{10, 6}
	interact(c)
}