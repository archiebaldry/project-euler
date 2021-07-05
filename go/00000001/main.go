package main

import "fmt"

func main() {
	fmt.Println(multiplesSum(-100))
	fmt.Println(multiplesSum(1000))
}

func multiplesSum(limit int) int {
	var sum int
	for i := 3; i < limit; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}
