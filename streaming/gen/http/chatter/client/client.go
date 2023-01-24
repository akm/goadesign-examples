// Code generated by goa v3.11.0, DO NOT EDIT.
//
// chatter client HTTP transport
//
// Command:
// $ goa gen goa.design/examples/streaming/design -o streaming

package client

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the chatter service endpoint HTTP clients.
type Client struct {
	// Login Doer is the HTTP client used to make requests to the login endpoint.
	LoginDoer goahttp.Doer

	// Echoer Doer is the HTTP client used to make requests to the echoer endpoint.
	EchoerDoer goahttp.Doer

	// Listener Doer is the HTTP client used to make requests to the listener
	// endpoint.
	ListenerDoer goahttp.Doer

	// Summary Doer is the HTTP client used to make requests to the summary
	// endpoint.
	SummaryDoer goahttp.Doer

	// Subscribe Doer is the HTTP client used to make requests to the subscribe
	// endpoint.
	SubscribeDoer goahttp.Doer

	// History Doer is the HTTP client used to make requests to the history
	// endpoint.
	HistoryDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme     string
	host       string
	encoder    func(*http.Request) goahttp.Encoder
	decoder    func(*http.Response) goahttp.Decoder
	dialer     goahttp.Dialer
	configurer *ConnConfigurer
}

// NewClient instantiates HTTP clients for all the chatter service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
	dialer goahttp.Dialer,
	cfn *ConnConfigurer,
) *Client {
	if cfn == nil {
		cfn = &ConnConfigurer{}
	}
	return &Client{
		LoginDoer:           doer,
		EchoerDoer:          doer,
		ListenerDoer:        doer,
		SummaryDoer:         doer,
		SubscribeDoer:       doer,
		HistoryDoer:         doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
		dialer:              dialer,
		configurer:          cfn,
	}
}

// Login returns an endpoint that makes HTTP requests to the chatter service
// login server.
func (c *Client) Login() goa.Endpoint {
	var (
		encodeRequest  = EncodeLoginRequest(c.encoder)
		decodeResponse = DecodeLoginResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildLoginRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.LoginDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("chatter", "login", err)
		}
		return decodeResponse(resp)
	}
}

// Echoer returns an endpoint that makes HTTP requests to the chatter service
// echoer server.
func (c *Client) Echoer() goa.Endpoint {
	var (
		encodeRequest  = EncodeEchoerRequest(c.encoder)
		decodeResponse = DecodeEchoerResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildEchoerRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		conn, resp, err := c.dialer.DialContext(ctx, req.URL.String(), req.Header)
		if err != nil {
			if resp != nil {
				return decodeResponse(resp)
			}
			return nil, goahttp.ErrRequestError("chatter", "echoer", err)
		}
		if c.configurer.EchoerFn != nil {
			conn = c.configurer.EchoerFn(conn, cancel)
		}
		stream := &EchoerClientStream{conn: conn}
		return stream, nil
	}
}

// Listener returns an endpoint that makes HTTP requests to the chatter service
// listener server.
func (c *Client) Listener() goa.Endpoint {
	var (
		encodeRequest  = EncodeListenerRequest(c.encoder)
		decodeResponse = DecodeListenerResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildListenerRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		conn, resp, err := c.dialer.DialContext(ctx, req.URL.String(), req.Header)
		if err != nil {
			if resp != nil {
				return decodeResponse(resp)
			}
			return nil, goahttp.ErrRequestError("chatter", "listener", err)
		}
		if c.configurer.ListenerFn != nil {
			conn = c.configurer.ListenerFn(conn, cancel)
		}
		stream := &ListenerClientStream{conn: conn}
		return stream, nil
	}
}

// Summary returns an endpoint that makes HTTP requests to the chatter service
// summary server.
func (c *Client) Summary() goa.Endpoint {
	var (
		encodeRequest  = EncodeSummaryRequest(c.encoder)
		decodeResponse = DecodeSummaryResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildSummaryRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		conn, resp, err := c.dialer.DialContext(ctx, req.URL.String(), req.Header)
		if err != nil {
			if resp != nil {
				return decodeResponse(resp)
			}
			return nil, goahttp.ErrRequestError("chatter", "summary", err)
		}
		if c.configurer.SummaryFn != nil {
			conn = c.configurer.SummaryFn(conn, cancel)
		}
		stream := &SummaryClientStream{conn: conn}
		return stream, nil
	}
}

// Subscribe returns an endpoint that makes HTTP requests to the chatter
// service subscribe server.
func (c *Client) Subscribe() goa.Endpoint {
	var (
		encodeRequest  = EncodeSubscribeRequest(c.encoder)
		decodeResponse = DecodeSubscribeResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildSubscribeRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		conn, resp, err := c.dialer.DialContext(ctx, req.URL.String(), req.Header)
		if err != nil {
			if resp != nil {
				return decodeResponse(resp)
			}
			return nil, goahttp.ErrRequestError("chatter", "subscribe", err)
		}
		if c.configurer.SubscribeFn != nil {
			conn = c.configurer.SubscribeFn(conn, cancel)
		}
		go func() {
			<-ctx.Done()
			conn.WriteControl(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseNormalClosure, "client closing connection"),
				time.Now().Add(time.Second),
			)
			conn.Close()
		}()
		stream := &SubscribeClientStream{conn: conn}
		return stream, nil
	}
}

// History returns an endpoint that makes HTTP requests to the chatter service
// history server.
func (c *Client) History() goa.Endpoint {
	var (
		encodeRequest  = EncodeHistoryRequest(c.encoder)
		decodeResponse = DecodeHistoryResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildHistoryRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		conn, resp, err := c.dialer.DialContext(ctx, req.URL.String(), req.Header)
		if err != nil {
			if resp != nil {
				return decodeResponse(resp)
			}
			return nil, goahttp.ErrRequestError("chatter", "history", err)
		}
		if c.configurer.HistoryFn != nil {
			conn = c.configurer.HistoryFn(conn, cancel)
		}
		go func() {
			<-ctx.Done()
			conn.WriteControl(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseNormalClosure, "client closing connection"),
				time.Now().Add(time.Second),
			)
			conn.Close()
		}()
		stream := &HistoryClientStream{conn: conn}
		view := resp.Header.Get("goa-view")
		stream.SetView(view)
		return stream, nil
	}
}
