package library

import "fmt"

type book struct {
	title  string
	isbn   string
	author string
}

func (b book) String() string {
	return fmt.Sprintf("book %s with isbn %v written by %s", b.title, b.isbn, b.author)
}
