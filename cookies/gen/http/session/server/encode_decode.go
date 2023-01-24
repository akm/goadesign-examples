// Code generated by goa v3.10.3, DO NOT EDIT.
//
// session HTTP server encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/cookies/design -o cookies

package server

import (
	"context"
	"io"
	"net/http"

	session "goa.design/examples/cookies/gen/session"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeCreateSessionResponse returns an encoder for responses returned by the
// session create_session endpoint.
func EncodeCreateSessionResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*session.CreateSessionResult)
		enc := encoder(ctx, w)
		body := NewCreateSessionResponseBody(res)
		sessionID := res.SessionID
		http.SetCookie(w, &http.Cookie{
			Name:   "SID",
			Value:  sessionID,
			MaxAge: 3600,
		})
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeCreateSessionRequest returns a decoder for requests sent to the
// session create_session endpoint.
func DecodeCreateSessionRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body CreateSessionRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateCreateSessionRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewCreateSessionPayload(&body)

		return payload, nil
	}
}

// EncodeUseSessionResponse returns an encoder for responses returned by the
// session use_session endpoint.
func EncodeUseSessionResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*session.UseSessionResult)
		enc := encoder(ctx, w)
		body := NewUseSessionResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeUseSessionRequest returns a decoder for requests sent to the session
// use_session endpoint.
func DecodeUseSessionRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			sessionID string
			err       error
			c         *http.Cookie
		)
		c, err = r.Cookie("SID")
		if err == http.ErrNoCookie {
			err = goa.MergeErrors(err, goa.MissingFieldError("SID", "cookie"))
		} else {
			sessionID = c.Value
		}
		if err != nil {
			return nil, err
		}
		payload := NewUseSessionPayload(sessionID)

		return payload, nil
	}
}
