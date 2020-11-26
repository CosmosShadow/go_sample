package main

import "fmt"

func main() {
	// init、position、post 可省略部分
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 切片
	strings := []string{"google", "runoob"}
	for i, s := range strings {
		fmt.Println(i, s)
	}
	numbers := [6]int{1, 2, 3, 5} 
	for i,x:= range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
	}

	// label标签，break的地方
	re:
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			break re
		}
	}
	// continue 也可以用label

	// 测试goto
	var a int = 10
	LOOP: for a < 20 {
		if a == 15 {
			/* 跳过迭代，继续执行for a < 20 */
			a++
			goto LOOP
		}
		fmt.Printf("a的值为 : %d\n", a)
		a++
	}
}














