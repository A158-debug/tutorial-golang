package main

import "fmt"

func someFunction(num string){
	fmt.Println("Hello " + num)
}

func main(){
	fmt.Println("Concurrency Patterns in Go")
	someFunction("Debug")
}