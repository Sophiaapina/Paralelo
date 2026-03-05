package main

import book "classwork16feb"

func main() {

	var myBook = book.Book{
		Title:  "los 5 lenguajes deo amor",
		Author: "Gary Chapman",
		Pages:  200,
	}

	myBook.PrintInfo()

	var anotherbook = book.NewBook("ensayo sobre la seguera", "jose saramago", 200)
}
