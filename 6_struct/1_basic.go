package main

import "fmt"

type Books struct {
	title string
	author string
	subject string
	book_id int
}

func main() {
	var Book1 Books        /* 声明 Book1 为 Books 类型 */
	var Book2 Books        /* 声明 Book2 为 Books 类型 */

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 1

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 2

	/* 打印 Book 信息 */
	for index, val := range [] Books {Book1, Book2} {
		fmt.Printf( "Book %d title : %s\n", index, val.title)
		fmt.Printf( "Book %d author : %s\n", index, val.author)
		fmt.Printf( "Book %d subject : %s\n", index, val.subject)
		fmt.Printf( "Book %d book_id : %d\n", index, val.book_id)
	}
}