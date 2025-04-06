//:= shorthand for declaring and initialzing variables only inside a function

package main

import "fmt"

func main(){
	message:= "Hello World"
	year:= 2001

	fmt.Println(message, year)
}