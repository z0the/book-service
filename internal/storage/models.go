package storage

import (
	"time"
)

func getAllModelsList() []interface{} {
	return []interface{}{
		&Book{}, &Author{},
	}
}

type Books []*Book

type Book struct {
	ID        string `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
	Name      string
	Length    uint
	Text      string
	Author
}

type Author struct {
	ID         string `gorm:"primarykey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `gorm:"index"`
	Name       string
	FamilyName string
	Age        uint
}

var anonymousAuthor = Author{
	ID:   "0",
	Name: "Anonym",
	Age:  0,
}
