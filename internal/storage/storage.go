package storage

import "context"

type Storage interface {
	CreateBook(ctx context.Context, book *Book) error
	GetBookByID(ctx context.Context, id string) (*Book, error)
}
