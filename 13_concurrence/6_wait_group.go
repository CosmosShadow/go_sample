package main

import (
	"fmt"
	"time"
	"sync"
)

// go是协程，在main退出是，所有协程都会退出，无论是否执行完成没有

var v = 0

func f() int {
	v++
	return v
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)		// 加入等待
	go func(s string, v int) {
		// defer: 压栈，routine退出是调用
		defer wg.Done()	// 执行完了，需要和上面的wg.Add对应数量，如果Add(2)，dead
		time.Sleep(time.Second * 1)
		fmt.Println(s, v)
	}("go task2: ", f())

	go fmt.Println("Main task1: ", f())

	wg.Wait()
	fmt.Println("exit")
}

// output
// go task1:  1
// Main task1:  3
// go task2:  2