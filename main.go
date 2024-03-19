package main

import (
	"fmt"
)

func divisibleBy3(n int) bool {
	return n%3 == 0
}

func divisibleBy5(n int) bool {
	return n%5 == 0
}

func main() {
	for i := 1; i <= 100; i++ {
		if divisibleBy3(i) {
			fmt.Println("Pin")
		}
		if divisibleBy5(i) {
			fmt.Println("Pan")
		}
	}
}
