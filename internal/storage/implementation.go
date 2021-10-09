package storage

import (
	"context"
	"time"

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
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return book, nil
}

func (s *storage) GetBookList(ctx context.Context, fromDate time.Time, limit int) (Books, error) {
	var books Books
	err := s.db.WithContext(ctx).Find(&books, "created_at >= ?", fromDate).Limit(limit).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return books, nil
}

func (s *storage) runMigrations() {
	err := s.db.AutoMigrate(getAllModelsList()...)
	if err != nil {
		panic(err)
	}
}
