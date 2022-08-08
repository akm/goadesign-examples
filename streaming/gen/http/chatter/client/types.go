// Code generated by goa v3.8.2, DO NOT EDIT.
//
// chatter HTTP client types
//
// Command:
// $ goa gen goa.design/examples/streaming/design -o streaming

package client

import (
	chatter "goa.design/examples/streaming/gen/chatter"
	chatterviews "goa.design/examples/streaming/gen/chatter/views"
	goa "goa.design/goa/v3/pkg"
)

// SummaryResponseBody is the type of the "chatter" service "summary" endpoint
// HTTP response body.
type SummaryResponseBody []*ChatSummaryResponse

// SubscribeResponseBody is the type of the "chatter" service "subscribe"
// endpoint HTTP response body.
type SubscribeResponseBody struct {
	// Message sent to the server
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	Action  *string `form:"action,omitempty" json:"action,omitempty" xml:"action,omitempty"`
	// Time at which the message was added
	AddedAt *string `form:"added_at,omitempty" json:"added_at,omitempty" xml:"added_at,omitempty"`
}

// HistoryResponseBody is the type of the "chatter" service "history" endpoint
// HTTP response body.
type HistoryResponseBody struct {
	// Message sent to the server
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Length of the message sent
	Length *int `form:"length,omitempty" json:"length,omitempty" xml:"length,omitempty"`
	// Time at which the message was sent
	SentAt *string `form:"sent_at,omitempty" json:"sent_at,omitempty" xml:"sent_at,omitempty"`
}

// ChatSummaryResponse is used to define fields on response body types.
type ChatSummaryResponse struct {
	// Message sent to the server
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Length of the message sent
	Length *int `form:"length,omitempty" json:"length,omitempty" xml:"length,omitempty"`
	// Time at which the message was sent
	SentAt *string `form:"sent_at,omitempty" json:"sent_at,omitempty" xml:"sent_at,omitempty"`
}

// NewLoginUnauthorized builds a chatter service login endpoint unauthorized
// error.
func NewLoginUnauthorized(body string) chatter.Unauthorized {
	v := chatter.Unauthorized(body)

	return v
}

// NewEchoerInvalidScopes builds a chatter service echoer endpoint
// invalid-scopes error.
func NewEchoerInvalidScopes(body string) chatter.InvalidScopes {
	v := chatter.InvalidScopes(body)

	return v
}

// NewEchoerUnauthorized builds a chatter service echoer endpoint unauthorized
// error.
func NewEchoerUnauthorized(body string) chatter.Unauthorized {
	v := chatter.Unauthorized(body)

	return v
}

// NewListenerInvalidScopes builds a chatter service listener endpoint
// invalid-scopes error.
func NewListenerInvalidScopes(body string) chatter.InvalidScopes {
	v := chatter.InvalidScopes(body)

	return v
}

// NewListenerUnauthorized builds a chatter service listener endpoint
// unauthorized error.
func NewListenerUnauthorized(body string) chatter.Unauthorized {
	v := chatter.Unauthorized(body)

	return v
}

// NewSummaryChatSummaryCollectionOK builds a "chatter" service "summary"
// endpoint result from a HTTP "OK" response.
func NewSummaryChatSummaryCollectionOK(body SummaryResponseBody) chatterviews.ChatSummaryCollectionView {
	v := make([]*chatterviews.ChatSummaryView, len(body))
	for i, val := range body {
		v[i] = unmarshalChatSummaryResponseToChatterviewsChatSummaryView(val)
	}

	return v
}

// NewSummaryInvalidScopes builds a chatter service summary endpoint
// invalid-scopes error.
func NewSummaryInvalidScopes(body string) chatter.InvalidScopes {
	v := chatter.InvalidScopes(body)

	return v
}

// NewSummaryUnauthorized builds a chatter service summary endpoint
// unauthorized error.
func NewSummaryUnauthorized(body string) chatter.Unauthorized {
	v := chatter.Unauthorized(body)

	return v
}

// NewSubscribeEventOK builds a "chatter" service "subscribe" endpoint result
// from a HTTP "OK" response.
func NewSubscribeEventOK(body *SubscribeResponseBody) *chatter.Event {
	v := &chatter.Event{
		Message: *body.Message,
		Action:  *body.Action,
		AddedAt: *body.AddedAt,
	}

	return v
}

// NewSubscribeInvalidScopes builds a chatter service subscribe endpoint
// invalid-scopes error.
func NewSubscribeInvalidScopes(body string) chatter.InvalidScopes {
	v := chatter.InvalidScopes(body)

	return v
}

// NewSubscribeUnauthorized builds a chatter service subscribe endpoint
// unauthorized error.
func NewSubscribeUnauthorized(body string) chatter.Unauthorized {
	v := chatter.Unauthorized(body)

	return v
}

// NewHistoryChatSummaryOK builds a "chatter" service "history" endpoint result
// from a HTTP "OK" response.
func NewHistoryChatSummaryOK(body *HistoryResponseBody) *chatterviews.ChatSummaryView {
	v := &chatterviews.ChatSummaryView{
		Message: body.Message,
		Length:  body.Length,
		SentAt:  body.SentAt,
	}

	return v
}

// NewHistoryInvalidScopes builds a chatter service history endpoint
// invalid-scopes error.
func NewHistoryInvalidScopes(body string) chatter.InvalidScopes {
	v := chatter.InvalidScopes(body)

	return v
}

// NewHistoryUnauthorized builds a chatter service history endpoint
// unauthorized error.
func NewHistoryUnauthorized(body string) chatter.Unauthorized {
	v := chatter.Unauthorized(body)

	return v
}

// ValidateSubscribeResponseBody runs the validations defined on
// SubscribeResponseBody
func ValidateSubscribeResponseBody(body *SubscribeResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Action == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("action", "body"))
	}
	if body.AddedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("added_at", "body"))
	}
	if body.Action != nil {
		if !(*body.Action == "added") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.action", *body.Action, []interface{}{"added"}))
		}
	}
	if body.AddedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.added_at", *body.AddedAt, goa.FormatDateTime))
	}
	return
}

// ValidateChatSummaryResponse runs the validations defined on
// ChatSummaryResponse
func ValidateChatSummaryResponse(body *ChatSummaryResponse) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.SentAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("sent_at", "body"))
	}
	if body.SentAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.sent_at", *body.SentAt, goa.FormatDateTime))
	}
	return
}
