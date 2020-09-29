package main

import (
	"fmt"
	"sync"
)

func main() {
	var ints []int
	var wg sync.WaitGroup
	var mx sync.Mutex
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mx.Lock()
			ints = append(ints, i)
			mx.Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Println(ints)
}