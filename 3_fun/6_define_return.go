package main

import "fmt"
 
func Test0() (ret string) {
	// 执行顺序3
	defer func() {
		err := recover()
		if err!= nil{
			ret=fmt.Sprint(err)
		}
	}()

	// 执行顺序2
	// panic也是一种defer，会压栈
	panic("this is a panic")

	// 执行顺序1
	return "normal"
}

func c() (i int) {
    defer func() { i++ }()
    return 1
}

func main() {
	fmt.Println(Test0())
	fmt.Println(c())
}