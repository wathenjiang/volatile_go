package main

import (
	"fmt"
	"time"
)

var flag int = 0

func main() {
	go func() {
		time.Sleep(1 * time.Second)
		flag = -1
		fmt.Println("flag was updated to -1")
	}()
	for flag != -1 {
	}
	fmt.Println("main goroutine read flag is -1 successful")
}
