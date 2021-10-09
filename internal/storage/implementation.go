package storage

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewStorage(logger *zap.SugaredLogger, db *gorm.DB) *storage {
	store := &storage{
		logger: logger,
		db:     db,
	}
	store.runMigrations()
	return store
}

type storage struct {
	logger *zap.SugaredLogger
	db     *gorm.DB
}

func (s *storage) CreateBook(ctx context.Context, book *Book) error {
	book.ID = uuid.NewString()
	return s.db.WithContext(ctx).Create(book).Error
}

func (s *storage) GetBookByID(ctx context.Context, id string) (*Book, error) {
	book := &Book{ID: id}
	err := s.db.WithContext(ctx).Take(book).Error
	if err != nil {
		return nil, err
	}
	return book, nil
	// return &Book{
	// 	ID:     id,
	// 	Name:   "Test Book",
	// 	Length: 15,
	// 	Text:   "text of the book",
	// 	Author: anonymousAuthor,
	// }, nil
}

func (s *storage) runMigrations() {
	err := s.db.AutoMigrate(getAllModelsList()...)
	if err != nil {
		panic(err)
	}
}
