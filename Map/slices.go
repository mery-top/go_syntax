//dynamic array

package main

import "fmt"

func main(){
	nums:= []int{1,2,3}
	nums = append(nums, 5)
	fmt.Println(nums)
}