package main

import (
	"fmt"
	"sync"
	"time"
)

/**
	Go routines are what go uses for its concurrency model and are based on the concept of coroutines (https://en.wikipedia.org/wiki/Coroutine)
	The main function runs as a goroutine as well, and it spawns other goroutines as required by the application.
	Once the main function exits, the execution ends regardless of whether or not other goroutines have finished executing.
	We can use channels/waitgroups to wait for the other goroutines

*/
func waitGroupGoRoutine(wg *sync.WaitGroup, order int){
	time.Sleep(time.Second * 2)
	fmt.Println("From wait group go routine:", order)
	wg.Done()
}


func channelGoRoutine(ch chan int, order int){
	time.Sleep(time.Second * 2)
	ch <- order
}


func UsingWaitGroup(){
	waitGroup := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		waitGroup.Add(1)
		// waitgroups should be always passed by reference
		go waitGroupGoRoutine(&waitGroup, i)
	}
	waitGroup.Wait()
}


func UsingChannels(){
	ch := make(chan int, 5)
	count := 0
	for i := 0; i < 5; i++ {
		go channelGoRoutine(ch, i)
	}

	// loop will run everytime a goroutine pushes to the ch channel
	for value := range ch {
		fmt.Println("From channel go routine:", value)
		count++
		if count == 5 {
			break
		}
	}
}

func main(){
	UsingWaitGroup()

	UsingChannels()
}