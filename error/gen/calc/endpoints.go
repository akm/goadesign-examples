// Code generated by goa v3.12.3, DO NOT EDIT.
//
// calc endpoints
//
// Command:
// $ goa gen goa.design/examples/error/design -o error

package calc

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "calc" service endpoints.
type Endpoints struct {
	Divide goa.Endpoint
	Add    goa.Endpoint
}

// NewEndpoints wraps the methods of the "calc" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Divide: NewDivideEndpoint(s),
		Add:    NewAddEndpoint(s),
	}
}

// Use applies the given middleware to all the "calc" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Divide = m(e.Divide)
	e.Add = m(e.Add)
}

// NewDivideEndpoint returns an endpoint function that calls the method
// "divide" of service "calc".
func NewDivideEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*DividePayload)
		return s.Divide(ctx, p)
	}
}

// NewAddEndpoint returns an endpoint function that calls the method "add" of
// service "calc".
func NewAddEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*AddPayload)
		return s.Add(ctx, p)
	}
}
