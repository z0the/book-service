package http

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"

	booksvc "books-service/internal/book_svc"
)

var typeAssertionError = errors.New("failed to do type assertion")

// Get book

type GetBookByIDReq struct {
	ID string `json:"id"`
}

type GetBookByIDResp struct {
	Book *booksvc.Book `json:"book,omitempty"`
}

func makeGetBookByID(c *Controller) endpoint.Endpoint {
	ep := func(ctx context.Context, request interface{}) (interface{}, error) {
		typedRequest, ok := request.(*GetBookByIDReq)
		if !ok {
			return nil, typeAssertionError
		}
		return c.getBookByID(ctx, typedRequest)
	}
	return ep
}

func (c *Controller) getBookByID(ctx context.Context, req *GetBookByIDReq) (*GetBookByIDResp, error) {
	book, err := c.bookSvc.GetBookByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	return &GetBookByIDResp{Book: book}, nil
}

// Create book

type Book struct {
	Name   string `json:"name"`
	Length uint   `json:"length"`
	Text   string `json:"text"`
}

func (b *Book) convertToService() *booksvc.Book {
	return &booksvc.Book{
		Name:   b.Name,
		Length: b.Length,
		Text:   b.Text,
	}
}

type Author struct {
	Name       string
	FamilyName string
	Age        uint
}

type CreateBookReq struct {
	Book `json:"book"`
}

type CreateBookResp struct{}

func makeCreateBook(c *Controller) endpoint.Endpoint {
	ep := func(ctx context.Context, request interface{}) (interface{}, error) {
		typedRequest, ok := request.(*CreateBookReq)
		if !ok {
			return nil, typeAssertionError
		}
		return c.createBook(ctx, typedRequest)
	}
	return ep
}

func (c *Controller) createBook(ctx context.Context, req *CreateBookReq) (*CreateBookResp, error) {
	err := c.bookSvc.CreateBook(ctx, req.Book.convertToService())
	if err != nil {
		return nil, err
	}
	return &CreateBookResp{}, nil
}
