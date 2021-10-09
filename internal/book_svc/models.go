package book_svc

import "books-service/internal/storage"

type Book struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Length uint   `json:"length"`
	Text   string `json:"text"`
	Author `json:"author"`
}

func newBooksFST(storeBooks storage.Books) Books {
	var res Books
	for _, storeBook := range storeBooks {
		res = append(res, newBookFST(storeBook))
	}
	return res
}

type Books []*Book

func (b *Book) convertToStorage() *storage.Book {
	return &storage.Book{
		ID:     b.ID,
		Name:   b.Name,
		Length: b.Length,
		Text:   b.Text,
		Author: *b.Author.toStorage(),
	}
}

// newBookFST - returns new Book from storage.Book
func newBookFST(b *storage.Book) *Book {
	if b == nil {
		return nil
	}
	return &Book{
		ID:     b.ID,
		Name:   b.Name,
		Length: b.Length,
		Text:   b.Text,
		Author: *newAuthorFST(&b.Author),
	}
}

type Author struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	FamilyName string `json:"family_name"`
	Age        uint   `json:"age"`
}

func (a *Author) toStorage() *storage.Author {
	return &storage.Author{
		ID:         a.ID,
		Name:       a.Name,
		FamilyName: a.FamilyName,
		Age:        a.Age,
	}
}

// newAuthorFST - returns new Book from storage.Book
func newAuthorFST(a *storage.Author) *Author {
	if a == nil {
		return nil
	}
	return &Author{
		ID:         a.ID,
		Name:       a.Name,
		FamilyName: a.FamilyName,
		Age:        a.Age,
	}
}
