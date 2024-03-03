package main

import "fmt"

func SumInts(m map[string]int64) int64 {
	var sum int64
	for _, v := range m{
		sum += v
	}
	return sum
}

func sumFloats(m map[string]float64) float64{
	var sum float64
	for _,v:= range m{
		sum += v
	}
	return sum
}

type Number interface{
	int64 | uint64 | float64
}

func SumIntsOrFloats[K string, T Number] (m map[string]T) T{
	var sum T
	for _, v := range m{
		sum += v
	}
	return sum
}

func main(){
	fmt.Println("Generics in Go")

	// initialize map fro the integer values
	intMap := map[string]int64{
		"first":34,
		"second": 56,
	}

	// initialize map for the float values
	floatMap := map[string]float64{
		"first": 34.5,	
		"second": 56.5,
	}

	fmt.Printf("Non Generic sums : %v and %v \n",SumInts(intMap),sumFloats(floatMap))
	fmt.Printf("Generic sums : %v and %v \n",SumIntsOrFloats(intMap),SumIntsOrFloats(floatMap))
	
}