// Code generated by goa v3.4.2, DO NOT EDIT.
//
// session HTTP client encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/cookies/design -o
// $(GOPATH)/src/goa.design/examples/cookies

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	session "goa.design/examples/cookies/gen/session"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// BuildCreateSessionRequest instantiates a HTTP request object with method and
// path set to call the "session" service "create_session" endpoint
func (c *Client) BuildCreateSessionRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateSessionSessionPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("session", "create_session", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateSessionRequest returns an encoder for requests sent to the
// session create_session server.
func EncodeCreateSessionRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*session.CreateSessionPayload)
		if !ok {
			return goahttp.ErrInvalidType("session", "create_session", "*session.CreateSessionPayload", v)
		}
		body := NewCreateSessionRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("session", "create_session", err)
		}
		return nil
	}
}

// DecodeCreateSessionResponse returns a decoder for responses returned by the
// session create_session endpoint. restoreBody controls whether the response
// body should be restored after having been read.
func DecodeCreateSessionResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body CreateSessionResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("session", "create_session", err)
			}
			err = ValidateCreateSessionResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("session", "create_session", err)
			}
			var (
				sessionID    string
				sessionIDRaw string

				cookies = resp.Cookies()
			)
			for _, c := range cookies {
				switch c.Name {
				case "SID":
					sessionIDRaw = c.Value
				}
			}
			if sessionIDRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("SID", "cookie"))
			}
			sessionID = sessionIDRaw
			if err != nil {
				return nil, goahttp.ErrValidationError("session", "create_session", err)
			}
			res := NewCreateSessionResultOK(&body, sessionID)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("session", "create_session", resp.StatusCode, string(body))
		}
	}
}

// BuildUseSessionRequest instantiates a HTTP request object with method and
// path set to call the "session" service "use_session" endpoint
func (c *Client) BuildUseSessionRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UseSessionSessionPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("session", "use_session", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUseSessionRequest returns an encoder for requests sent to the session
// use_session server.
func EncodeUseSessionRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*session.UseSessionPayload)
		if !ok {
			return goahttp.ErrInvalidType("session", "use_session", "*session.UseSessionPayload", v)
		}
		{
			v := p.SessionID
			req.AddCookie(&http.Cookie{
				Name:  "SID",
				Value: v,
			})
		}
		return nil
	}
}

// DecodeUseSessionResponse returns a decoder for responses returned by the
// session use_session endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeUseSessionResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body UseSessionResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("session", "use_session", err)
			}
			err = ValidateUseSessionResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("session", "use_session", err)
			}
			res := NewUseSessionResultOK(&body)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("session", "use_session", resp.StatusCode, string(body))
		}
	}
}
