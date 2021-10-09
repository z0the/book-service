package book_svc

import "context"

// Service - interface of service
type Service interface {
	CreateBook(ctx context.Context, book *Book) error
	GetBookByID(ctx context.Context, id string) (*Book, error)
	GetBookList(ctx context.Context, fromID uint, limit int) (Books, error)
}
type CreateBookReq struct {
	Book `json:"book"`
}

type CreateBookResp struct {
	Err string `json:"err,omitempty"`
}
