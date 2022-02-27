package main

import (
	"fmt"
	"time"
)

// var flag int = 0
var c chan int = make(chan int)

func main() {
	go func() {
		time.Sleep(1 * time.Second)
		c <- -1
		fmt.Println("child goroutine send to channel a -1")
	}()
	<-c
	fmt.Println("main goroutine read from channel with a -1")
}
