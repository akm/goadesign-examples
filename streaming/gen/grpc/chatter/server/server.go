// Code generated by goa v3.9.0, DO NOT EDIT.
//
// chatter gRPC server
//
// Command:
// $ goa gen goa.design/examples/streaming/design -o streaming

package server

import (
	"context"
	"errors"

	chatter "goa.design/examples/streaming/gen/chatter"
	chatterpb "goa.design/examples/streaming/gen/grpc/chatter/pb"
	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc/codes"
)

// Server implements the chatterpb.ChatterServer interface.
type Server struct {
	LoginH     goagrpc.UnaryHandler
	EchoerH    goagrpc.StreamHandler
	ListenerH  goagrpc.StreamHandler
	SummaryH   goagrpc.StreamHandler
	SubscribeH goagrpc.StreamHandler
	HistoryH   goagrpc.StreamHandler
	chatterpb.UnimplementedChatterServer
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the expr.
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
type ErrorNamer interface {
	ErrorName() string
}

// GoaErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type GoaErrorNamer interface {
	GoaErrorName() string
}

type adaptErrorNamer struct {
	ErrorNamer
}

func (err adaptErrorNamer) GoaErrorName() string {
	return err.ErrorName()
}

func (err adaptErrorNamer) Error() string {
	return err.ErrorNamer.(error).Error()
}

// EchoerServerStream implements the chatter.EchoerServerStream interface.
type EchoerServerStream struct {
	stream chatterpb.Chatter_EchoerServer
}

// ListenerServerStream implements the chatter.ListenerServerStream interface.
type ListenerServerStream struct {
	stream chatterpb.Chatter_ListenerServer
}

// SummaryServerStream implements the chatter.SummaryServerStream interface.
type SummaryServerStream struct {
	stream chatterpb.Chatter_SummaryServer
	view   string
}

// SubscribeServerStream implements the chatter.SubscribeServerStream interface.
type SubscribeServerStream struct {
	stream chatterpb.Chatter_SubscribeServer
}

// HistoryServerStream implements the chatter.HistoryServerStream interface.
type HistoryServerStream struct {
	stream chatterpb.Chatter_HistoryServer
	view   string
}

// New instantiates the server struct with the chatter service endpoints.
func New(e *chatter.Endpoints, uh goagrpc.UnaryHandler, sh goagrpc.StreamHandler) *Server {
	return &Server{
		LoginH:     NewLoginHandler(e.Login, uh),
		EchoerH:    NewEchoerHandler(e.Echoer, sh),
		ListenerH:  NewListenerHandler(e.Listener, sh),
		SummaryH:   NewSummaryHandler(e.Summary, sh),
		SubscribeH: NewSubscribeHandler(e.Subscribe, sh),
		HistoryH:   NewHistoryHandler(e.History, sh),
	}
}

// NewLoginHandler creates a gRPC handler which serves the "chatter" service
// "login" endpoint.
func NewLoginHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeLoginRequest, EncodeLoginResponse)
	}
	return h
}

// Login implements the "Login" method in chatterpb.ChatterServer interface.
func (s *Server) Login(ctx context.Context, message *chatterpb.LoginRequest) (*chatterpb.LoginResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "login")
	ctx = context.WithValue(ctx, goa.ServiceKey, "chatter")
	resp, err := s.LoginH.Handle(ctx, message)
	if err != nil {
		var deprecatedErrorNamer ErrorNamer
		if errors.As(err, &deprecatedErrorNamer) {
			err = adaptErrorNamer{deprecatedErrorNamer}
		}
		var en GoaErrorNamer
		if errors.As(err, &en) {
			switch en.GoaErrorName() {
			case "unauthorized":
				return nil, goagrpc.NewStatusError(codes.Unauthenticated, err, goagrpc.NewErrorResponse(err))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*chatterpb.LoginResponse), nil
}

// NewEchoerHandler creates a gRPC handler which serves the "chatter" service
// "echoer" endpoint.
func NewEchoerHandler(endpoint goa.Endpoint, h goagrpc.StreamHandler) goagrpc.StreamHandler {
	if h == nil {
		h = goagrpc.NewStreamHandler(endpoint, DecodeEchoerRequest)
	}
	return h
}

// Echoer implements the "Echoer" method in chatterpb.ChatterServer interface.
func (s *Server) Echoer(stream chatterpb.Chatter_EchoerServer) error {
	ctx := stream.Context()
	ctx = context.WithValue(ctx, goa.MethodKey, "echoer")
	ctx = context.WithValue(ctx, goa.ServiceKey, "chatter")
	p, err := s.EchoerH.Decode(ctx, nil)
	if err != nil {
		var deprecatedErrorNamer ErrorNamer
		if errors.As(err, &deprecatedErrorNamer) {
			err = adaptErrorNamer{deprecatedErrorNamer}
		}
		var en GoaErrorNamer
		if errors.As(err, &en) {
			switch en.GoaErrorName() {
			case "unauthorized":
				return goagrpc.NewStatusError(codes.Unauthenticated, err, goagrpc.NewErrorResponse(err))
			case "invalid-scopes":
				return goagrpc.NewStatusError(codes.Unauthenticated, err, goagrpc.NewErrorResponse(err))
			}
		}
		return goagrpc.EncodeError(err)
	}
	ep := &chatter.EchoerEndpointInput{
		Stream:  &EchoerServerStream{stream: stream},
		Payload: p.(*chatter.EchoerPayload),
	}
	err = s.EchoerH.Handle(ctx, ep)
	if err != nil {
		var deprecatedErrorNamer ErrorNamer
		if errors.As(err, &deprecatedErrorNamer) {
			err = adaptErrorNamer{deprecatedErrorNamer}
		}
		var en GoaErrorNamer
		if errors.As(err, &en) {
			switch en.GoaErrorName() {
			case "unauthorized":
				return goagrpc.NewStatusError(codes.Unauthenticated, err, goagrpc.NewErrorResponse(err))
			case "invalid-scopes":
				return goagrpc.NewStatusError(codes.Unauthenticated, err, goagrpc.NewErrorResponse(err))
			}
		}
		return goagrpc.EncodeError(err)
	}
	return nil
}

// NewListenerHandler creates a gRPC handler which serves the "chatter" service
// "listener" endpoint.
func NewListenerHandler(endpoint goa.Endpoint, h goagrpc.StreamHandler) goagrpc.StreamHandler {
	if h == nil {
		h = goagrpc.NewStreamHandler(endpoint, DecodeListenerRequest)
	}
	return h
}

// Listener implements the "Listener" method in chatterpb.ChatterServer
// interface.
func (s *Server) Listener(stream chatterpb.Chatter_ListenerServer) error {
	ctx := stream.Context()
	ctx = context.WithValue(ctx, goa.MethodKey, "listener")
	ctx = context.WithValue(ctx, goa.ServiceKey, "chatter")
	p, err := s.ListenerH.Decode(ctx, nil)
	if err != nil {
		var deprecatedErrorNamer ErrorNamer
		if errors.As(err, &deprecatedErrorNamer) {
			err = adaptErrorNamer{deprecatedErrorNamer}
		}
		var en GoaErrorNamer
		if errors.As(err, &en) {
			switch en.GoaErrorName() {
			case "unauthorized":
				return goagrpc.NewStatusError(codes.Unauthenticated, err, goagrpc.NewErrorResponse(err))
			case "invalid-scopes":
				return goagrpc.NewStatusError(codes.Unauthenticated, err, goagrpc.NewErrorResponse(err))
			}
		}
		return goagrpc.EncodeError(err)
	}
	ep := &chatter.ListenerEndpointInput{
		Stream:  &ListenerServerStream{stream: stream},
		Payload: p.(*chatter.ListenerPayload),
	}
	err = s.ListenerH.Handle(ctx, ep)
	if err != nil {
		var deprecatedErrorNamer ErrorNamer
		if errors.As(err, &deprecatedErrorNamer) {
			err = adaptErrorNamer{deprecatedErrorNamer}
		}
		var en GoaErrorNamer
		if errors.As(err, &en) {
			switch en.GoaErrorName() {
			case "unauthorized":
				return goagrpc.NewStatusError(codes.Unauthenticated, err, goagrpc.NewErrorResponse(err))
			case "invalid-scopes":
				return goagrpc.NewStatusError(codes.Unauthenticated, err, goagrpc.NewErrorResponse(err))
			}
		}
		return goagrpc.EncodeError(err)
	}
	return nil
}

// NewSummaryHandler creates a gRPC handler which serves the "chatter" service
// "summary" endpoint.
func NewSummaryHandler(endpoint goa.Endpoint, h goagrpc.StreamHandler) goagrpc.StreamHandler {
	if h == nil {
		h = goagrpc.NewStreamHandler(endpoint, DecodeSummaryRequest)
	}
	return h
}

// Summary implements the "Summary" method in chatterpb.ChatterServer interface.
func (s *Server) Summary(stream chatterpb.Chatter_SummaryServer) error {
	ctx := stream.Context()
	ctx = context.WithValue(ctx, goa.MethodKey, "summary")
	ctx = context.WithValue(ctx, goa.ServiceKey, "chatter")
	p, err := s.SummaryH.Decode(ctx, nil)
	if err != nil {
		var deprecatedErrorNamer ErrorNamer
		if errors.As(err, &deprecatedErrorNamer) {
			err = adaptErrorNamer{deprecatedErrorNamer}
		}
		var en GoaErrorNamer
		if errors.As(err, &en) {
			switch en.GoaErrorName() {
			case "unauthorized":
				return goagrpc.NewStatusError(codes.Unauthenticated, err, goagrpc.NewErrorResponse(err))
			case "invalid-scopes":
				return goagrpc.NewStatusError(codes.Unauthenticated, err, goagrpc.NewErrorResponse(err))
			}
		}
		return goagrpc.EncodeError(err)
	}
	ep := &chatter.SummaryEndpointInput{
		Stream:  &SummaryServerStream{stream: stream},
		Payload: p.(*chatter.SummaryPayload),
	}
	err = s.SummaryH.Handle(ctx, ep)
	if err != nil {
		var deprecatedErrorNamer ErrorNamer
		if errors.As(err, &deprecatedErrorNamer) {
			err = adaptErrorNamer{deprecatedErrorNamer}
		}
		var en GoaErrorNamer
		if errors.As(err, &en) {
			switch en.GoaErrorName() {
			case "unauthorized":
				return goagrpc.NewStatusError(codes.Unauthenticated, err, goagrpc.NewErrorResponse(err))
			case "invalid-scopes":
				return goagrpc.NewStatusError(codes.Unauthenticated, err, goagrpc.NewErrorResponse(err))
			}
		}
		return goagrpc.EncodeError(err)
	}
	return nil
}

// NewSubscribeHandler creates a gRPC handler which serves the "chatter"
// service "subscribe" endpoint.
func NewSubscribeHandler(endpoint goa.Endpoint, h goagrpc.StreamHandler) goagrpc.StreamHandler {
	if h == nil {
		h = goagrpc.NewStreamHandler(endpoint, DecodeSubscribeRequest)
	}
	return h
}

// Subscribe implements the "Subscribe" method in chatterpb.ChatterServer
// interface.
func (s *Server) Subscribe(message *chatterpb.SubscribeRequest, stream chatterpb.Chatter_SubscribeServer) error {
	ctx := stream.Context()
	ctx = context.WithValue(ctx, goa.MethodKey, "subscribe")
	ctx = context.WithValue(ctx, goa.ServiceKey, "chatter")
	p, err := s.SubscribeH.Decode(ctx, message)
	if err != nil {
		var deprecatedErrorNamer ErrorNamer
		if errors.As(err, &deprecatedErrorNamer) {
			err = adaptErrorNamer{deprecatedErrorNamer}
		}
		var en GoaErrorNamer
		if errors.As(err, &en) {
			switch en.GoaErrorName() {
			case "unauthorized":
				return goagrpc.NewStatusError(codes.Unauthenticated, err, goagrpc.NewErrorResponse(err))
			case "invalid-scopes":
				return goagrpc.NewStatusError(codes.Unauthenticated, err, goagrpc.NewErrorResponse(err))
			}
		}
		return goagrpc.EncodeError(err)
	}
	ep := &chatter.SubscribeEndpointInput{
		Stream:  &SubscribeServerStream{stream: stream},
		Payload: p.(*chatter.SubscribePayload),
	}
	err = s.SubscribeH.Handle(ctx, ep)
	if err != nil {
		var deprecatedErrorNamer ErrorNamer
		if errors.As(err, &deprecatedErrorNamer) {
			err = adaptErrorNamer{deprecatedErrorNamer}
		}
		var en GoaErrorNamer
		if errors.As(err, &en) {
			switch en.GoaErrorName() {
			case "unauthorized":
				return goagrpc.NewStatusError(codes.Unauthenticated, err, goagrpc.NewErrorResponse(err))
			case "invalid-scopes":
				return goagrpc.NewStatusError(codes.Unauthenticated, err, goagrpc.NewErrorResponse(err))
			}
		}
		return goagrpc.EncodeError(err)
	}
	return nil
}

// NewHistoryHandler creates a gRPC handler which serves the "chatter" service
// "history" endpoint.
func NewHistoryHandler(endpoint goa.Endpoint, h goagrpc.StreamHandler) goagrpc.StreamHandler {
	if h == nil {
		h = goagrpc.NewStreamHandler(endpoint, DecodeHistoryRequest)
	}
	return h
}

// History implements the "History" method in chatterpb.ChatterServer interface.
func (s *Server) History(message *chatterpb.HistoryRequest, stream chatterpb.Chatter_HistoryServer) error {
	ctx := stream.Context()
	ctx = context.WithValue(ctx, goa.MethodKey, "history")
	ctx = context.WithValue(ctx, goa.ServiceKey, "chatter")
	p, err := s.HistoryH.Decode(ctx, message)
	if err != nil {
		var deprecatedErrorNamer ErrorNamer
		if errors.As(err, &deprecatedErrorNamer) {
			err = adaptErrorNamer{deprecatedErrorNamer}
		}
		var en GoaErrorNamer
		if errors.As(err, &en) {
			switch en.GoaErrorName() {
			case "unauthorized":
				return goagrpc.NewStatusError(codes.Unauthenticated, err, goagrpc.NewErrorResponse(err))
			case "invalid-scopes":
				return goagrpc.NewStatusError(codes.Unauthenticated, err, goagrpc.NewErrorResponse(err))
			}
		}
		return goagrpc.EncodeError(err)
	}
	ep := &chatter.HistoryEndpointInput{
		Stream:  &HistoryServerStream{stream: stream},
		Payload: p.(*chatter.HistoryPayload),
	}
	err = s.HistoryH.Handle(ctx, ep)
	if err != nil {
		var deprecatedErrorNamer ErrorNamer
		if errors.As(err, &deprecatedErrorNamer) {
			err = adaptErrorNamer{deprecatedErrorNamer}
		}
		var en GoaErrorNamer
		if errors.As(err, &en) {
			switch en.GoaErrorName() {
			case "unauthorized":
				return goagrpc.NewStatusError(codes.Unauthenticated, err, goagrpc.NewErrorResponse(err))
			case "invalid-scopes":
				return goagrpc.NewStatusError(codes.Unauthenticated, err, goagrpc.NewErrorResponse(err))
			}
		}
		return goagrpc.EncodeError(err)
	}
	return nil
}

// Send streams instances of "chatterpb.EchoerResponse" to the "echoer"
// endpoint gRPC stream.
func (s *EchoerServerStream) Send(res string) error {
	v := NewProtoEchoerResponse(res)
	return s.stream.Send(v)
}

// Recv reads instances of "chatterpb.EchoerStreamingRequest" from the "echoer"
// endpoint gRPC stream.
func (s *EchoerServerStream) Recv() (string, error) {
	var res string
	v, err := s.stream.Recv()
	if err != nil {
		return res, err
	}
	return NewEchoerStreamingRequest(v), nil
}

func (s *EchoerServerStream) Close() error {
	// nothing to do here
	return nil
}

// Recv reads instances of "chatterpb.ListenerStreamingRequest" from the
// "listener" endpoint gRPC stream.
func (s *ListenerServerStream) Recv() (string, error) {
	var res string
	v, err := s.stream.Recv()
	if err != nil {
		return res, err
	}
	return NewListenerStreamingRequest(v), nil
}

func (s *ListenerServerStream) Close() error {
	// synchronize stream
	return s.stream.SendAndClose(&chatterpb.ListenerResponse{})
}

// SendAndClose streams instances of "chatterpb.ChatSummaryCollection" to the
// "summary" endpoint gRPC stream.
func (s *SummaryServerStream) SendAndClose(res chatter.ChatSummaryCollection) error {
	vres := chatter.NewViewedChatSummaryCollection(res, "default")
	v := NewProtoChatSummaryCollection(vres.Projected)
	return s.stream.SendAndClose(v)
}

// Recv reads instances of "chatterpb.SummaryStreamingRequest" from the
// "summary" endpoint gRPC stream.
func (s *SummaryServerStream) Recv() (string, error) {
	var res string
	v, err := s.stream.Recv()
	if err != nil {
		return res, err
	}
	return NewSummaryStreamingRequest(v), nil
}

// Send streams instances of "chatterpb.SubscribeResponse" to the "subscribe"
// endpoint gRPC stream.
func (s *SubscribeServerStream) Send(res *chatter.Event) error {
	v := NewProtoSubscribeResponse(res)
	return s.stream.Send(v)
}

func (s *SubscribeServerStream) Close() error {
	// nothing to do here
	return nil
}

// Send streams instances of "chatterpb.HistoryResponse" to the "history"
// endpoint gRPC stream.
func (s *HistoryServerStream) Send(res *chatter.ChatSummary) error {
	vres := chatter.NewViewedChatSummary(res, s.view)
	v := NewProtoHistoryResponse(vres.Projected)
	return s.stream.Send(v)
}

func (s *HistoryServerStream) Close() error {
	// nothing to do here
	return nil
}

// SetView sets the view.
func (s *HistoryServerStream) SetView(view string) {
	s.view = view
}
