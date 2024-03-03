package main

import (
	"fmt"
	"unicode/utf8"
)

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var result int = numerator / denominator
	var remainder int = numerator % denominator

	var err error
	if denominator == 0 {
		return 0, 0, err
	}
	return result, remainder,err // default value of err = nil
}

func main() {

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)
	fmt.Println(intDivision(intNum1, intNum2))

	var myString string = "Hello My name is Debug"
	printMe(myString)
	fmt.Println(myString)
	fmt.Println("Length of string is ", utf8.RuneCountInString(myString))

	var1, var2 := 1, 2
	println(var1, var2)

	// default value of all int is = 0
	// default value of bool is false
	// default value of string is ""

    intArr := [3]int32{1, 2, 3}
	fmt.Println(intArr)

	intArr2 := []int32{1, 2, 3, 4, 5}
	fmt.Println(intArr2)

	intArr2 = append(intArr2, 6)
	fmt.Println(intArr2)

	var intSlice []int32 = []int32{7,8}
	var intSlice2 []int32 = []int32{9,10}

	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

}
