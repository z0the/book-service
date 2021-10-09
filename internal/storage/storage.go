package storage

import (
	"context"
	"time"
)

type Storage interface {
	CreateBook(ctx context.Context, book *Book) error
	GetBookByID(ctx context.Context, id string) (*Book, error)
	GetBookList(ctx context.Context, fromDate time.Time, limit int) (Books, error)
}
