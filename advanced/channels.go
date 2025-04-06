package main

import(
	"fmt"
)

func main(){
	ch1 := make(chan string)
	ch2:= make(chan string)

	go func(){
		ch1 <- "One: Message is sent from c1"
	}()

	go func(){
		ch2 <- "Two: Message is received in c2"
	}()

	select{
	case msg1 := <-ch1:
		fmt.Println(msg1)

	case msg2:= <- ch2:
		fmt.Println(msg2)
	}




}