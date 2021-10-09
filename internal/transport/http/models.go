package http

import booksvc "books-service/internal/book_svc"

type Book struct {
	Name   string `json:"name"`
	Length uint   `json:"length"`
	Text   string `json:"text"`
	Author `json:"author"`
}

func (b *Book) convertToService() *booksvc.Book {
	return &booksvc.Book{
		Name:   b.Name,
		Length: b.Length,
		Text:   b.Text,
		Author: *b.Author.convertToService(),
	}
}

type Author struct {
	Name       string
	FamilyName string
	Age        uint
}

func (a *Author) convertToService() *booksvc.Author {
	return &booksvc.Author{
		Name:       a.Name,
		FamilyName: a.FamilyName,
		Age:        a.Age,
	}
}
