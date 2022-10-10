package main

import "fmt"

// golang 没有 class ，但更强调类型 type，实现 class 只是 type 功能的一部分
type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	fmt.Println(Books{"Go", "me", "Go", 123})

	var Book1 Books

	Book1.title = "Go"
	Book1.author = "me"
	Book1.subject = "Go"
	Book1.book_id = 123

	printBook(&Book1)
}

func printBook(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}
