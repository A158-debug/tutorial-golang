package main

import "fmt"

func main(){
	var p* int32 = new(int32)
	fmt.Println(p, *p)

	*p = 100
	fmt.Println(p, *p)

	slice := []int32{1,2,3}
	sliceCopy := slice
	fmt.Println(slice, sliceCopy) 

	sliceCopy[2] = 4
	fmt.Println(slice, sliceCopy)

	

}