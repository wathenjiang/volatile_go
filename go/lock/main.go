package main

import (
	"fmt"
	"sync"
	"time"
)

var flag int = 0
var lock sync.Mutex

func main() {
	go func() {
		time.Sleep(1 * time.Second)
		lock.Lock()
		defer lock.Unlock()
		flag = -1
		fmt.Println("flag was updated to -1")
	}()
	for {
		lock.Lock()
		if flag == -1 {
			break
		}
		lock.Unlock()
	}
	fmt.Println("main goroutine read flag is -1 successful")
}
