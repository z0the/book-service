package http

import (
	"net/http"

	"books-service/internal/book_svc"
)

func NewController(bookSvc book_svc.Service) *Controller {
	ctrl := &Controller{
		bookSvc: bookSvc,
	}
	ctrl.init()
	return ctrl
}

type Controller struct {
	bookSvc             book_svc.Service
	sugaredEndpointList []SugaredEndpoint
}

func (c *Controller) init() {
	c.sugaredEndpointList = []SugaredEndpoint{
		{
			Path:      "/getBookByID",
			Method:    http.MethodPost,
			RequestT:  GetBookByIDReq{},
			ResponseT: GetBookByIDResp{},
			Endpoint:  makeGetBookByIDEndpoint(c),
		},
		{
			Path:      "/createBook",
			Method:    http.MethodPost,
			RequestT:  CreateBookReq{},
			ResponseT: CreateBookResp{},
			Endpoint:  makeCreateBookEndpoint(c),
		},
		{
			Path:      "/getBookList",
			Method:    http.MethodPost,
			RequestT:  GetBookListReq{},
			ResponseT: GetBookListResp{},
			Endpoint:  makeGetBookListEndpoint(c),
		},
	}
}

func (c *Controller) GetSugaredEndpointList() []SugaredEndpoint {
	return c.sugaredEndpointList
}
