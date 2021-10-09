// Package sugarep add a functionality around go-kit.Endpoint
package sugarep

import (
	"github.com/go-kit/kit/endpoint"
)

type EndpointDescription struct {
	Name            string
	Path            string
	Endpoint        endpoint.Endpoint
	Request         interface{}
	Response        interface{}
	RequestDecoder  DecodeFunc
	ResponseEncoder EncodeFunc
	Description     string
	// Middleware      []Middleware
	// Options         []ServerOption
	Methods  string
	Settings interface{}
}

type DecodeFunc func()
type EncodeFunc func()
