// Code generated by goa v3.13.0, DO NOT EDIT.
//
// chatter gRPC client types
//
// Command:
// $ goa gen goa.design/examples/streaming/design -o streaming

package client

import (
	chatter "goa.design/examples/streaming/gen/chatter"
	chatterviews "goa.design/examples/streaming/gen/chatter/views"
	chatterpb "goa.design/examples/streaming/gen/grpc/chatter/pb"
	goa "goa.design/goa/v3/pkg"
)

// NewProtoLoginRequest builds the gRPC request type from the payload of the
// "login" endpoint of the "chatter" service.
func NewProtoLoginRequest() *chatterpb.LoginRequest {
	message := &chatterpb.LoginRequest{}
	return message
}

// NewLoginResult builds the result type of the "login" endpoint of the
// "chatter" service from the gRPC response type.
func NewLoginResult(message *chatterpb.LoginResponse) string {
	result := message.Field
	return result
}

func NewEchoerResponseEchoerResponse(v *chatterpb.EchoerResponse) string {
	result := v.Field
	return result
}

func NewProtoEchoerStreamingRequest(spayload string) *chatterpb.EchoerStreamingRequest {
	v := &chatterpb.EchoerStreamingRequest{}
	v.Field = spayload
	return v
}

func NewProtoListenerStreamingRequest(spayload string) *chatterpb.ListenerStreamingRequest {
	v := &chatterpb.ListenerStreamingRequest{}
	v.Field = spayload
	return v
}

func NewChatSummaryCollectionChatSummaryCollection(v *chatterpb.ChatSummaryCollection) chatterviews.ChatSummaryCollectionView {
	vresult := make([]*chatterviews.ChatSummaryView, len(v.Field))
	for i, val := range v.Field {
		vresult[i] = &chatterviews.ChatSummaryView{
			Message: &val.Message_,
			SentAt:  &val.SentAt,
		}
		if val.Length != nil {
			length := int(*val.Length)
			vresult[i].Length = &length
		}
	}
	return vresult
}

func NewProtoSummaryStreamingRequest(spayload string) *chatterpb.SummaryStreamingRequest {
	v := &chatterpb.SummaryStreamingRequest{}
	v.Field = spayload
	return v
}

// NewProtoSubscribeRequest builds the gRPC request type from the payload of
// the "subscribe" endpoint of the "chatter" service.
func NewProtoSubscribeRequest() *chatterpb.SubscribeRequest {
	message := &chatterpb.SubscribeRequest{}
	return message
}

func NewSubscribeResponseEvent(v *chatterpb.SubscribeResponse) *chatter.Event {
	result := &chatter.Event{
		Message: v.Message_,
		Action:  v.Action,
		AddedAt: v.AddedAt,
	}
	return result
}

// NewProtoHistoryRequest builds the gRPC request type from the payload of the
// "history" endpoint of the "chatter" service.
func NewProtoHistoryRequest() *chatterpb.HistoryRequest {
	message := &chatterpb.HistoryRequest{}
	return message
}

func NewHistoryResponseChatSummaryView(v *chatterpb.HistoryResponse) *chatterviews.ChatSummaryView {
	vresult := &chatterviews.ChatSummaryView{
		Message: &v.Message_,
		SentAt:  &v.SentAt,
	}
	if v.Length != nil {
		length := int(*v.Length)
		vresult.Length = &length
	}
	return vresult
}

// ValidateChatSummaryCollection runs the validations defined on
// ChatSummaryCollection.
func ValidateChatSummaryCollection(stream *chatterpb.ChatSummaryCollection) (err error) {
	for _, e := range stream.Field {
		if e != nil {
			if err2 := ValidateChatSummary(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateChatSummary runs the validations defined on ChatSummary.
func ValidateChatSummary(elem *chatterpb.ChatSummary) (err error) {
	err = goa.MergeErrors(err, goa.ValidateFormat("elem.sent_at", elem.SentAt, goa.FormatDateTime))
	return
}

// ValidateSubscribeResponse runs the validations defined on SubscribeResponse.
func ValidateSubscribeResponse(stream *chatterpb.SubscribeResponse) (err error) {
	if !(stream.Action == "added") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError("stream.action", stream.Action, []any{"added"}))
	}
	err = goa.MergeErrors(err, goa.ValidateFormat("stream.added_at", stream.AddedAt, goa.FormatDateTime))
	return
}

// ValidateHistoryResponse runs the validations defined on HistoryResponse.
func ValidateHistoryResponse(stream *chatterpb.HistoryResponse) (err error) {
	err = goa.MergeErrors(err, goa.ValidateFormat("stream.sent_at", stream.SentAt, goa.FormatDateTime))
	return
}
