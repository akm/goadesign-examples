// Code generated by goa v3.0.0, DO NOT EDIT.
//
// secured_service HTTP server encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/security/design -o
// $(GOPATH)/src/goa.design/examples/security

package server

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	securedservice "goa.design/examples/security/gen/secured_service"
	goa "goa.design/goa/v3"
	goahttp "goa.design/goa/v3/http"
)

// EncodeSigninResponse returns an encoder for responses returned by the
// secured_service signin endpoint.
func EncodeSigninResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*securedservice.Creds)
		enc := encoder(ctx, w)
		body := NewSigninResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeSigninRequest returns a decoder for requests sent to the
// secured_service signin endpoint.
func DecodeSigninRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		payload := NewSigninPayload()
		user, pass, ok := r.BasicAuth()
		if !ok {
			return nil, goa.MissingFieldError("Authorization", "header")
		}
		payload.Username = user
		payload.Password = pass

		return payload, nil
	}
}

// EncodeSigninError returns an encoder for errors returned by the signin
// secured_service endpoint.
func EncodeSigninError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unauthorized":
			res := v.(securedservice.Unauthorized)
			enc := encoder(ctx, w)
			body := NewSigninUnauthorizedResponseBody(res)
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeSecureResponse returns an encoder for responses returned by the
// secured_service secure endpoint.
func EncodeSecureResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(string)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeSecureRequest returns a decoder for requests sent to the
// secured_service secure endpoint.
func DecodeSecureRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			fail  *bool
			token string
			err   error
		)
		{
			failRaw := r.URL.Query().Get("fail")
			if failRaw != "" {
				v, err2 := strconv.ParseBool(failRaw)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("fail", failRaw, "boolean"))
				}
				fail = &v
			}
		}
		token = r.Header.Get("Authorization")
		if token == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Authorization", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewSecurePayload(fail, token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}

		return payload, nil
	}
}

// EncodeSecureError returns an encoder for errors returned by the secure
// secured_service endpoint.
func EncodeSecureError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "invalid-scopes":
			res := v.(securedservice.InvalidScopes)
			enc := encoder(ctx, w)
			body := NewSecureInvalidScopesResponseBody(res)
			w.Header().Set("goa-error", "invalid-scopes")
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "unauthorized":
			res := v.(securedservice.Unauthorized)
			enc := encoder(ctx, w)
			body := NewSecureUnauthorizedResponseBody(res)
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeDoublySecureResponse returns an encoder for responses returned by the
// secured_service doubly_secure endpoint.
func EncodeDoublySecureResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(string)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeDoublySecureRequest returns a decoder for requests sent to the
// secured_service doubly_secure endpoint.
func DecodeDoublySecureRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			key   string
			token string
			err   error
		)
		key = r.URL.Query().Get("k")
		if key == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("k", "query string"))
		}
		token = r.Header.Get("Authorization")
		if token == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Authorization", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewDoublySecurePayload(key, token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}

		return payload, nil
	}
}

// EncodeDoublySecureError returns an encoder for errors returned by the
// doubly_secure secured_service endpoint.
func EncodeDoublySecureError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "invalid-scopes":
			res := v.(securedservice.InvalidScopes)
			enc := encoder(ctx, w)
			body := NewDoublySecureInvalidScopesResponseBody(res)
			w.Header().Set("goa-error", "invalid-scopes")
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "unauthorized":
			res := v.(securedservice.Unauthorized)
			enc := encoder(ctx, w)
			body := NewDoublySecureUnauthorizedResponseBody(res)
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeAlsoDoublySecureResponse returns an encoder for responses returned by
// the secured_service also_doubly_secure endpoint.
func EncodeAlsoDoublySecureResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(string)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeAlsoDoublySecureRequest returns a decoder for requests sent to the
// secured_service also_doubly_secure endpoint.
func DecodeAlsoDoublySecureRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			key        *string
			oauthToken *string
			token      *string
		)
		keyRaw := r.URL.Query().Get("k")
		if keyRaw != "" {
			key = &keyRaw
		}
		oauthTokenRaw := r.URL.Query().Get("oauth")
		if oauthTokenRaw != "" {
			oauthToken = &oauthTokenRaw
		}
		tokenRaw := r.Header.Get("X-Authorization")
		if tokenRaw != "" {
			token = &tokenRaw
		}
		payload := NewAlsoDoublySecurePayload(key, oauthToken, token)
		user, pass, _ := r.BasicAuth()
		payload.Username = &user
		payload.Password = &pass
		if payload.Token != nil {
			if strings.Contains(*payload.Token, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Token, " ", 2)[1]
				payload.Token = &cred
			}
		}

		return payload, nil
	}
}

// EncodeAlsoDoublySecureError returns an encoder for errors returned by the
// also_doubly_secure secured_service endpoint.
func EncodeAlsoDoublySecureError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "invalid-scopes":
			res := v.(securedservice.InvalidScopes)
			enc := encoder(ctx, w)
			body := NewAlsoDoublySecureInvalidScopesResponseBody(res)
			w.Header().Set("goa-error", "invalid-scopes")
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "unauthorized":
			res := v.(securedservice.Unauthorized)
			enc := encoder(ctx, w)
			body := NewAlsoDoublySecureUnauthorizedResponseBody(res)
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}
