// Code generated by goa v3.7.10, DO NOT EDIT.
//
// calc client
//
// Command:
// $ goa gen goa.design/examples/basic/design -o basic

package calc

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "calc" service client.
type Client struct {
	MultiplyEndpoint goa.Endpoint
}

// NewClient initializes a "calc" service client given the endpoints.
func NewClient(multiply goa.Endpoint) *Client {
	return &Client{
		MultiplyEndpoint: multiply,
	}
}

// Multiply calls the "multiply" endpoint of the "calc" service.
func (c *Client) Multiply(ctx context.Context, p *MultiplyPayload) (res int, err error) {
	var ires interface{}
	ires, err = c.MultiplyEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(int), nil
}
