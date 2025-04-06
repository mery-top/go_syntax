package main

import "fmt"

type Animal interface{
	Walker
	Talker
}

type Walker interface{
	walk(name string)
}

type Talker interface{
	talk(sound string)
}

type Dog struct{
	name string
	sound string
}


func (d Dog) walk(name string){
	fmt.Println(name, "can walk")
} 

func (d Dog) talk(sound string){
	fmt.Println("The Sound is", sound)
}

func interact(a Animal, name string, sound string){
	a.walk(name)
	a.talk(sound)
}

func main(){
	d:=Dog{"Dog", "vow"}
	interact(d, d.name, d.sound)
}