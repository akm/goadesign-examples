// Code generated by goa v3.9.0, DO NOT EDIT.
//
// swagger endpoints
//
// Command:
// $ goa gen goa.design/examples/cellar/design -o cellar

package swagger

import (
	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "swagger" service endpoints.
type Endpoints struct {
}

// NewEndpoints wraps the methods of the "swagger" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{}
}

// Use applies the given middleware to all the "swagger" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
}
