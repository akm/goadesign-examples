// Code generated by goa v3.10.0, DO NOT EDIT.
//
// calc client HTTP transport
//
// Command:
// $ goa gen goa.design/examples/basic/design -o basic

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the calc service endpoint HTTP clients.
type Client struct {
	// Multiply Doer is the HTTP client used to make requests to the multiply
	// endpoint.
	MultiplyDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the calc service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		MultiplyDoer:        doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Multiply returns an endpoint that makes HTTP requests to the calc service
// multiply server.
func (c *Client) Multiply() goa.Endpoint {
	var (
		decodeResponse = DecodeMultiplyResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildMultiplyRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.MultiplyDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("calc", "multiply", err)
		}
		return decodeResponse(resp)
	}
}
