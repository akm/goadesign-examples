// Code generated by goa v3.10.2, DO NOT EDIT.
//
// session HTTP client types
//
// Command:
// $ goa gen goa.design/examples/cookies/design -o cookies

package client

import (
	session "goa.design/examples/cookies/gen/session"
	goa "goa.design/goa/v3/pkg"
)

// CreateSessionRequestBody is the type of the "session" service
// "create_session" endpoint HTTP request body.
type CreateSessionRequestBody struct {
	// Name of session
	Name string `form:"name" json:"name" xml:"name"`
}

// CreateSessionResponseBody is the type of the "session" service
// "create_session" endpoint HTTP response body.
type CreateSessionResponseBody struct {
	// User message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// UseSessionResponseBody is the type of the "session" service "use_session"
// endpoint HTTP response body.
type UseSessionResponseBody struct {
	// User message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// NewCreateSessionRequestBody builds the HTTP request body from the payload of
// the "create_session" endpoint of the "session" service.
func NewCreateSessionRequestBody(p *session.CreateSessionPayload) *CreateSessionRequestBody {
	body := &CreateSessionRequestBody{
		Name: p.Name,
	}
	return body
}

// NewCreateSessionResultOK builds a "session" service "create_session"
// endpoint result from a HTTP "OK" response.
func NewCreateSessionResultOK(body *CreateSessionResponseBody, sessionID string) *session.CreateSessionResult {
	v := &session.CreateSessionResult{
		Message: *body.Message,
	}
	v.SessionID = sessionID

	return v
}

// NewUseSessionResultOK builds a "session" service "use_session" endpoint
// result from a HTTP "OK" response.
func NewUseSessionResultOK(body *UseSessionResponseBody) *session.UseSessionResult {
	v := &session.UseSessionResult{
		Message: *body.Message,
	}

	return v
}

// ValidateCreateSessionResponseBody runs the validations defined on
// create_session_response_body
func ValidateCreateSessionResponseBody(body *CreateSessionResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateUseSessionResponseBody runs the validations defined on
// use_session_response_body
func ValidateUseSessionResponseBody(body *UseSessionResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}
