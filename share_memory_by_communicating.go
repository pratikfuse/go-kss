package main

import "fmt"

func WriteOnly(channel chan<- int, order int) {
	channel <- order
}



func main() {
	var ints []int

	channel := make(chan int, 10)

	for i := 0; i < 10; i++ {
		go WriteOnly(channel, i)
	}

	for i := range channel {
		ints = append(ints, i)

		if len(ints) == 10 {
			break
		}
	}

	fmt.Printf("Ints %v", ints)
}
