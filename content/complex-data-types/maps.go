package main

import "fmt"

func main() {
	var f map[string]int
	f = make(map[string]int)
	fmt.Println(f)

	fruits := map[string]int{
		"apple":  2,
		"orange": 3,
		"pear":   1,
	}
	fmt.Println(fruits)
	for fruit, count := range fruits {
		fmt.Println("Fruit:", fruit, "Count:", count)
	}
}
