package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	timeout := time.After(time.Second * 2)
	go func(c chan int) {
		defer fmt.Println("go end")
		// select用来做超时处理
		select {
			case str1 := <- c:
				fmt.Println("Receive ", str1, time.Now())
			case <- timeout:
				fmt.Println("time out", time.Now())
		}
	}(c)

	time.Sleep(time.Second * 5)
	fmt.Println("Main exit", time.Now())
}