package main

import (
	"fmt"
	"time"
)

/*
	select allows to listen on multiple channels at once
	*/


func Run( ch chan <- string, waitFor time.Duration){
	time.Sleep(waitFor )
	// close the channel on the production end; i.e. the go routine that pushes to the routine
	ch <- "go routine 1"
}


func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)


	go Run(ch1, time.Second * 10)
	go Run(ch2, time.Second * 5)

	select{
	case <-ch1:
		fmt.Println("recieved from channel 1")
	case <-ch2:
		fmt.Println("recieved from channel 2")
	}
}
