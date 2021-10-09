package book_svc

import (
	"context"

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
	if id == "12" {
		return nil, terrors.MakeBadRequestErr("wrong id")
	}
	book, err := s.storage.GetBookByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if book == nil {
		return nil, terrors.MakeNotFoundErr("nil book")
	}
	return newBookFST(book), nil
}

func (s *service) GetBookList(ctx context.Context, fromID uint, limit int) (Books, error) {
	return nil, nil
}
