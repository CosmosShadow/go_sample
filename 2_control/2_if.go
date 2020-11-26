package main

import "fmt"

func main() {
	var a = 10
	if a < 20 {
		fmt.Printf("%d is less than 20\n", a)	// 注意Printf和Println有区别
	} else {
		fmt.Printf("%d >= 20\n", a)
	}
	fmt.Printf("a is %d\n", a)

	if a == 25 {
		fmt.Println("true")
	} else if a < 25 {
		fmt.Println("too small")
	} else {
		fmt.Println("too big")
	}
}