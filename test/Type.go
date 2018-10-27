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
	Book1.title = "Go "
	Book1.author = "om"
	Book1.subject = "G"
	Book1.book_id = 64

	/* book 2 描述 */
	Book2.title = "P"
	Book2.author = "ww"
	Book2.subject = "Pytho"
	Book2.book_id = 6495700


	// 创建一个新的结构体
	fmt.Println(Books{"Go 语言", "om", "Go", 6495407})

	// 也可以使用 key => value 格式
	fmt.Println(Books{title: "Go", author: "ww", subject: "Go", book_id: 6495407})

	// 忽略的字段为 0 或 空
	fmt.Println(Books{title: "Go", author: "m"})
}