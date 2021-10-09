package http

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"

	"github.com/go-kit/kit/endpoint"
	kitHttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"books-service/internal/transport/terrors"
)

type SugaredEndpoint struct {
	Path                string
	Method              string
	MiddlewareList      []endpoint.Middleware
	RequestT, ResponseT interface{}
	endpoint.Endpoint
}

type SugaredEndpointProvider interface {
	GetSugaredEndpointList() []SugaredEndpoint
}

func LoadHTTPRouterWithEndpoints(router *mux.Router, epProvider SugaredEndpointProvider) {
	for _, sugaredEp := range epProvider.GetSugaredEndpointList() {
		router.Handle(
			sugaredEp.Path,
			kitHttp.NewServer(
				sugaredEp.Endpoint,
				getDecodeJSONRequestFunc(sugaredEp.RequestT),
				EncodeJSONResponse,
				kitHttp.ServerErrorEncoder(terrors.GetErrEncoderFunc(log.Default())))).
			Methods(sugaredEp.Method)
	}
}

type ResponseBody struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
}

func EncodeJSONResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	result := ResponseBody{
		Success: true,
		Data:    response,
	}
	return kitHttp.EncodeJSONResponse(ctx, w, result)
}

func EmptyRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	var req interface{}
	return req, nil
}

func getDecodeJSONRequestFunc(RequestT interface{}) kitHttp.DecodeRequestFunc {
	if RequestT == nil {
		return EmptyRequest
	}
	return func(_ context.Context, r *http.Request) (i interface{}, err error) {
		var body []byte

		body, err = ioutil.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}

		prototype, err := getNewPtrToType(RequestT)
		if err != nil {
			return nil, err
		}
		if len(body) == 0 {
			return prototype, nil
		}

		err = json.Unmarshal(body, prototype)
		if err != nil {
			return nil, err
		}

		return prototype, nil
	}

}

// getNewPtrToType - receives an object of T type, allocates memory and returns a pointer *T to type
// Returns err if obj is a pointer *T to type
func getNewPtrToType(obj interface{}) (interface{}, error) {
	if reflect.TypeOf(obj).Kind() == reflect.Ptr {
		return nil, errors.New("wrong object kind")
	}
	return reflect.New(reflect.TypeOf(obj)).Interface(), nil
}
