// Code generated by goa v3.10.0, DO NOT EDIT.
//
// api_key_service client
//
// Command:
// $ goa gen goa.design/examples/security/hierarchy/design -o security/hierarchy

package apikeyservice

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "api_key_service" service client.
type Client struct {
	DefaultEndpoint  goa.Endpoint
	SecureEndpoint   goa.Endpoint
	UnsecureEndpoint goa.Endpoint
}

// NewClient initializes a "api_key_service" service client given the endpoints.
func NewClient(default_, secure, unsecure goa.Endpoint) *Client {
	return &Client{
		DefaultEndpoint:  default_,
		SecureEndpoint:   secure,
		UnsecureEndpoint: unsecure,
	}
}

// Default calls the "default" endpoint of the "api_key_service" service.
func (c *Client) Default(ctx context.Context, p *DefaultPayload) (err error) {
	_, err = c.DefaultEndpoint(ctx, p)
	return
}

// Secure calls the "secure" endpoint of the "api_key_service" service.
func (c *Client) Secure(ctx context.Context, p *SecurePayload) (err error) {
	_, err = c.SecureEndpoint(ctx, p)
	return
}

// Unsecure calls the "unsecure" endpoint of the "api_key_service" service.
func (c *Client) Unsecure(ctx context.Context) (err error) {
	_, err = c.UnsecureEndpoint(ctx, nil)
	return
}
