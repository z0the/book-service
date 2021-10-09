package http

import (
	"context"
	"errors"
	"time"

	"github.com/go-kit/kit/endpoint"

	booksvc "books-service/internal/book_svc"
)

var wrongBodyError = errors.New("wrong request body")

// getBook

type GetBookByIDReq struct {
	ID string `json:"id"`
}

type GetBookByIDResp struct {
	Book *booksvc.Book `json:"book,omitempty"`
}

func makeGetBookByIDEndpoint(c *Controller) endpoint.Endpoint {
	ep := func(ctx context.Context, request interface{}) (interface{}, error) {
		typedRequest, ok := request.(*GetBookByIDReq)
		if !ok {
			return nil, wrongBodyError
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

// createBook

type CreateBookReq struct {
	Book `json:"book"`
}

type CreateBookResp struct{}

func makeCreateBookEndpoint(c *Controller) endpoint.Endpoint {
	ep := func(ctx context.Context, request interface{}) (interface{}, error) {
		typedRequest, ok := request.(*CreateBookReq)
		if !ok {
			return nil, wrongBodyError
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

// getBookList

type GetBookListReq struct {
	FromDateUnix int64 `json:"from_date"`
	Limit        int   `json:"limit"`
}

type GetBookListResp struct {
	booksvc.Books `json:"books"`
}

func makeGetBookListEndpoint(c *Controller) endpoint.Endpoint {
	ep := func(ctx context.Context, request interface{}) (interface{}, error) {
		typedRequest, ok := request.(*GetBookListReq)
		if !ok {
			return nil, wrongBodyError
		}
		return c.getBookList(ctx, typedRequest)
	}
	return ep
}

func (c *Controller) getBookList(ctx context.Context, req *GetBookListReq) (*GetBookListResp, error) {
	books, err := c.bookSvc.GetBookList(ctx, time.Unix(req.FromDateUnix, 0), req.Limit)
	if err != nil {
		return nil, err
	}
	return &GetBookListResp{Books: books}, nil
}
