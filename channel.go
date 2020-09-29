package main

import (
	"fmt"
	"strconv"
	"time"
)


func Fetch(ch chan <- string, message string)  {
	// delay for one second
	time.Sleep(time.Second * 1)
	ch <- message
}

func main() {
	unBufferedChannel := make(chan string)

	bufferedChannel := make(chan string, 10)
	go Fetch(unBufferedChannel, "from unbuffered channel")

	bufferedChannelCount := 0

	for i:= 0; i < 10; i++ {
		go Fetch(bufferedChannel, "from buffered channel: " + strconv.Itoa(i))
	}

	// reading from an unbuffered channel
	response := <- unBufferedChannel

	fmt.Println(response)

	// will listen on bufferedChannel until its empty
	for v := range bufferedChannel {
		fmt.Println(v)
		bufferedChannelCount++

		// close the channel once all 10 goroutines have finished execution
		if bufferedChannelCount ==  10 {
			close(bufferedChannel)
		}

	}

}
