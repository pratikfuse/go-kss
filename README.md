# GO Programming Language

![Go lang](https://miro.medium.com/max/700/1*w6xcm-WXHTTpSv037EoshA.jpeg)

## Installation
1. Head over to https://golang.org/doc/install
2. Follow instructions for your operating system

### [Go Modules](https://blog.golang.org/using-go-modules)
Collection of Packages in a go application and is defined by a go.mod file at the root.
Modules are created using the command:

    go mod init <module_name>
   

## Go routines
Go routines are what go uses for its concurrency model. The main function can also be thought
of as the (main) Go routine which in turn spawns other go routines in the application.


## Channels
Channels are what go routines use to communicate and pass messages between each other.

### creating a channel
Channel are created with the following syntax:

    channel := make(chan int, 10)

First variable takes a chan keyword and type of data (int in this case) that is passed over the channel. The data type can be base
types such as int, float, string or user defined struct types.

Second parameter is optional and defines a buffer which is the count of values that can be stored in 
a channel until a go routine reads from that channel. A channel can accept values from goroutine as long
as there is room in buffer for that channel.



 # Useful Links

 [Golang blog](https://blog.golang.org/)  
 [Go concepts](https://kuree.gitbooks.io/the-go-programming-language-report/)  
 [Go Proverbs](https://go-proverbs.github.io/)  
 [Goroutines](https://medium.com/technofullel/understanding-golang-and-goroutines-72ac3c9a014d)  
 [How docker is built in go](https://www.youtube.com/watch?v=Utf-A4rODH8)  
 
 
    

    

 