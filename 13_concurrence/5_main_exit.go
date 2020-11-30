package main

import (
	"fmt"
	"time"
)

// go是协程，在main退出时，所有协程都会退出，无论是否执行完成没有

var v = 0

func f() int {
	v++
	return v
}

func main() {
	go fmt.Println("go task1: ", f())
	go func(s string, v int) {
		time.Sleep(time.Second * 1)
		fmt.Println(s, v)
	}("go task2: ", f())

	go fmt.Println("Main task1: ", f())

	time.Sleep(time.Second * 3)
}

// output
// go task1:  1
// Main task1:  3
// go task2:  2