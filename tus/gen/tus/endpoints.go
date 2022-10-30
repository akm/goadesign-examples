// Code generated by goa v3.10.2, DO NOT EDIT.
//
// tus endpoints
//
// Command:
// $ goa gen goa.design/examples/tus/design -o tus

package tus

import (
	"context"
	"io"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "tus" service endpoints.
type Endpoints struct {
	Head    goa.Endpoint
	Patch   goa.Endpoint
	Options goa.Endpoint
	Post    goa.Endpoint
	Delete  goa.Endpoint
}

// PatchRequestData holds both the payload and the HTTP request body reader of
// the "patch" method.
type PatchRequestData struct {
	// Payload is the method payload.
	Payload *PatchPayload
	// Body streams the HTTP request body.
	Body io.ReadCloser
}

// PostRequestData holds both the payload and the HTTP request body reader of
// the "post" method.
type PostRequestData struct {
	// Payload is the method payload.
	Payload *PostPayload
	// Body streams the HTTP request body.
	Body io.ReadCloser
}

// NewEndpoints wraps the methods of the "tus" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Head:    NewHeadEndpoint(s),
		Patch:   NewPatchEndpoint(s),
		Options: NewOptionsEndpoint(s),
		Post:    NewPostEndpoint(s),
		Delete:  NewDeleteEndpoint(s),
	}
}

// Use applies the given middleware to all the "tus" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Head = m(e.Head)
	e.Patch = m(e.Patch)
	e.Options = m(e.Options)
	e.Post = m(e.Post)
	e.Delete = m(e.Delete)
}

// NewHeadEndpoint returns an endpoint function that calls the method "head" of
// service "tus".
func NewHeadEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*HeadPayload)
		return s.Head(ctx, p)
	}
}

// NewPatchEndpoint returns an endpoint function that calls the method "patch"
// of service "tus".
func NewPatchEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		ep := req.(*PatchRequestData)
		return s.Patch(ctx, ep.Payload, ep.Body)
	}
}

// NewOptionsEndpoint returns an endpoint function that calls the method
// "options" of service "tus".
func NewOptionsEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Options(ctx)
	}
}

// NewPostEndpoint returns an endpoint function that calls the method "post" of
// service "tus".
func NewPostEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		ep := req.(*PostRequestData)
		return s.Post(ctx, ep.Payload, ep.Body)
	}
}

// NewDeleteEndpoint returns an endpoint function that calls the method
// "delete" of service "tus".
func NewDeleteEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*DeletePayload)
		return s.Delete(ctx, p)
	}
}
