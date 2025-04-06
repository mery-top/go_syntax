package main

import (
	"fmt"
	"time"
)


func printMsg(){
	fmt.Println("Hello from Meerthika")
}

func main(){
	go printMsg()
	time.Sleep(1* time.Second)
}