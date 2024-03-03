package main

import "fmt"

func main() {
	fmt.Println("Maps in Go")

	// To initialize a map, use the built in make function
	m := make(map[string]int)
	fmt.Println(m)

	m["routes"] = 1
	m["age"] = 22
	m["distance"] = 34
	fmt.Println(m)

	i := m["routes"]
	fmt.Println(i)

	// If the requested key doesn’t exist, we get the value type’s zero value.
	j := m["root"]
	fmt.Println(j)

	length_of_map := len(m)
	fmt.Println(length_of_map)

	// The built in delete function removes an entry from the map:
	delete(m, "routes")
	fmt.Println(m)

	value_assign, bool_value := m["routes"]
	fmt.Println(value_assign, bool_value)

	// In this statement, the first value (i) is assigned the value stored under the key "route". If that key doesn’t exist, i is the value type’s zero value (0). The second value (ok) is a bool that is true if the key exists in the map, and false if not

	_, ok := m["age"]
	fmt.Println(ok)

	// To iterate over the key/value pairs of a map, use the range form of the for loop.

	for key, value := range m {
		fmt.Println(key, value)
	}

	// To initialize a map with some data, use a map literal:

	commits := map[string]int{
		"rsc": 3711,
		"r":   2138,
		"gri": 1908,
		"adg": 912,
	}
	fmt.Println(commits)

}
