package book_svc

import (
	"context"
	"time"

	"go.uber.org/zap"

	"books-service/internal/storage"
	"books-service/internal/transport/terrors"
)

type ServiceConfig struct {
	Logger  *zap.SugaredLogger
	Storage storage.Storage
}

func NewService(cfg *ServiceConfig) *service {
	return &service{
		logger:  cfg.Logger,
		storage: cfg.Storage,
	}
}

// service - implementation of interface Service
type service struct {
	logger  *zap.SugaredLogger
	storage storage.Storage
}

func (s *service) CreateBook(ctx context.Context, book *Book) error {
	return s.storage.CreateBook(ctx, book.convertToStorage())
}

func (s *service) GetBookByID(ctx context.Context, id string) (*Book, error) {
	book, err := s.storage.GetBookByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if book == nil {
		return nil, terrors.MakeNotFoundErr("nil book")
	}
	return newBookFST(book), nil
}

func (s *service) GetBookList(ctx context.Context, fromDate time.Time, limit int) (Books, error) {
	if limit > 100 {
		limit = 100
	}
	books, err := s.storage.GetBookList(ctx, fromDate, limit)
	if err != nil {
		return nil, err
	}
	if books == nil || len(books) == 0 {
		return nil, terrors.MakeNotFoundErr("no books")
	}
	return newBooksFST(books), nil
}
