// Code generated by goa v3.13.0, DO NOT EDIT.
//
// chatter client
//
// Command:
// $ goa gen goa.design/examples/streaming/design -o streaming

package chatter

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "chatter" service client.
type Client struct {
	LoginEndpoint     goa.Endpoint
	EchoerEndpoint    goa.Endpoint
	ListenerEndpoint  goa.Endpoint
	SummaryEndpoint   goa.Endpoint
	SubscribeEndpoint goa.Endpoint
	HistoryEndpoint   goa.Endpoint
}

// NewClient initializes a "chatter" service client given the endpoints.
func NewClient(login, echoer, listener, summary, subscribe, history goa.Endpoint) *Client {
	return &Client{
		LoginEndpoint:     login,
		EchoerEndpoint:    echoer,
		ListenerEndpoint:  listener,
		SummaryEndpoint:   summary,
		SubscribeEndpoint: subscribe,
		HistoryEndpoint:   history,
	}
}

// Login calls the "login" endpoint of the "chatter" service.
// Login may return the following errors:
//   - "unauthorized" (type Unauthorized)
//   - error: internal error
func (c *Client) Login(ctx context.Context, p *LoginPayload) (res string, err error) {
	var ires any
	ires, err = c.LoginEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}

// Echoer calls the "echoer" endpoint of the "chatter" service.
// Echoer may return the following errors:
//   - "unauthorized" (type Unauthorized)
//   - "invalid-scopes" (type InvalidScopes)
//   - error: internal error
func (c *Client) Echoer(ctx context.Context, p *EchoerPayload) (res EchoerClientStream, err error) {
	var ires any
	ires, err = c.EchoerEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(EchoerClientStream), nil
}

// Listener calls the "listener" endpoint of the "chatter" service.
// Listener may return the following errors:
//   - "unauthorized" (type Unauthorized)
//   - "invalid-scopes" (type InvalidScopes)
//   - error: internal error
func (c *Client) Listener(ctx context.Context, p *ListenerPayload) (res ListenerClientStream, err error) {
	var ires any
	ires, err = c.ListenerEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(ListenerClientStream), nil
}

// Summary calls the "summary" endpoint of the "chatter" service.
// Summary may return the following errors:
//   - "unauthorized" (type Unauthorized)
//   - "invalid-scopes" (type InvalidScopes)
//   - error: internal error
func (c *Client) Summary(ctx context.Context, p *SummaryPayload) (res SummaryClientStream, err error) {
	var ires any
	ires, err = c.SummaryEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(SummaryClientStream), nil
}

// Subscribe calls the "subscribe" endpoint of the "chatter" service.
// Subscribe may return the following errors:
//   - "unauthorized" (type Unauthorized)
//   - "invalid-scopes" (type InvalidScopes)
//   - error: internal error
func (c *Client) Subscribe(ctx context.Context, p *SubscribePayload) (res SubscribeClientStream, err error) {
	var ires any
	ires, err = c.SubscribeEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(SubscribeClientStream), nil
}

// History calls the "history" endpoint of the "chatter" service.
// History may return the following errors:
//   - "unauthorized" (type Unauthorized)
//   - "invalid-scopes" (type InvalidScopes)
//   - error: internal error
func (c *Client) History(ctx context.Context, p *HistoryPayload) (res HistoryClientStream, err error) {
	var ires any
	ires, err = c.HistoryEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(HistoryClientStream), nil
}
