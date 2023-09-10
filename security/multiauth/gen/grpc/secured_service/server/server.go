// Code generated by goa v3.13.0, DO NOT EDIT.
//
// secured_service gRPC server
//
// Command:
// $ goa gen goa.design/examples/security/multiauth/design -o security/multiauth

package server

import (
	"context"
	"errors"

	secured_servicepb "goa.design/examples/security/multiauth/gen/grpc/secured_service/pb"
	securedservice "goa.design/examples/security/multiauth/gen/secured_service"
	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc/codes"
)

// Server implements the secured_servicepb.SecuredServiceServer interface.
type Server struct {
	SigninH           goagrpc.UnaryHandler
	SecureH           goagrpc.UnaryHandler
	DoublySecureH     goagrpc.UnaryHandler
	AlsoDoublySecureH goagrpc.UnaryHandler
	secured_servicepb.UnimplementedSecuredServiceServer
}

// New instantiates the server struct with the secured_service service
// endpoints.
func New(e *securedservice.Endpoints, uh goagrpc.UnaryHandler) *Server {
	return &Server{
		SigninH:           NewSigninHandler(e.Signin, uh),
		SecureH:           NewSecureHandler(e.Secure, uh),
		DoublySecureH:     NewDoublySecureHandler(e.DoublySecure, uh),
		AlsoDoublySecureH: NewAlsoDoublySecureHandler(e.AlsoDoublySecure, uh),
	}
}

// NewSigninHandler creates a gRPC handler which serves the "secured_service"
// service "signin" endpoint.
func NewSigninHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeSigninRequest, EncodeSigninResponse)
	}
	return h
}

// Signin implements the "Signin" method in
// secured_servicepb.SecuredServiceServer interface.
func (s *Server) Signin(ctx context.Context, message *secured_servicepb.SigninRequest) (*secured_servicepb.SigninResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "signin")
	ctx = context.WithValue(ctx, goa.ServiceKey, "secured_service")
	resp, err := s.SigninH.Handle(ctx, message)
	if err != nil {
		var en goa.GoaErrorNamer
		if errors.As(err, &en) {
			switch en.GoaErrorName() {
			case "unauthorized":
				return nil, goagrpc.NewStatusError(codes.Unauthenticated, err, goagrpc.NewErrorResponse(err))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*secured_servicepb.SigninResponse), nil
}

// NewSecureHandler creates a gRPC handler which serves the "secured_service"
// service "secure" endpoint.
func NewSecureHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeSecureRequest, EncodeSecureResponse)
	}
	return h
}

// Secure implements the "Secure" method in
// secured_servicepb.SecuredServiceServer interface.
func (s *Server) Secure(ctx context.Context, message *secured_servicepb.SecureRequest) (*secured_servicepb.SecureResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "secure")
	ctx = context.WithValue(ctx, goa.ServiceKey, "secured_service")
	resp, err := s.SecureH.Handle(ctx, message)
	if err != nil {
		var en goa.GoaErrorNamer
		if errors.As(err, &en) {
			switch en.GoaErrorName() {
			case "invalid-scopes":
				return nil, goagrpc.NewStatusError(codes.Unauthenticated, err, goagrpc.NewErrorResponse(err))
			case "unauthorized":
				return nil, goagrpc.NewStatusError(codes.Unauthenticated, err, goagrpc.NewErrorResponse(err))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*secured_servicepb.SecureResponse), nil
}

// NewDoublySecureHandler creates a gRPC handler which serves the
// "secured_service" service "doubly_secure" endpoint.
func NewDoublySecureHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeDoublySecureRequest, EncodeDoublySecureResponse)
	}
	return h
}

// DoublySecure implements the "DoublySecure" method in
// secured_servicepb.SecuredServiceServer interface.
func (s *Server) DoublySecure(ctx context.Context, message *secured_servicepb.DoublySecureRequest) (*secured_servicepb.DoublySecureResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "doubly_secure")
	ctx = context.WithValue(ctx, goa.ServiceKey, "secured_service")
	resp, err := s.DoublySecureH.Handle(ctx, message)
	if err != nil {
		var en goa.GoaErrorNamer
		if errors.As(err, &en) {
			switch en.GoaErrorName() {
			case "invalid-scopes":
				return nil, goagrpc.NewStatusError(codes.Unauthenticated, err, goagrpc.NewErrorResponse(err))
			case "unauthorized":
				return nil, goagrpc.NewStatusError(codes.Unauthenticated, err, goagrpc.NewErrorResponse(err))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*secured_servicepb.DoublySecureResponse), nil
}

// NewAlsoDoublySecureHandler creates a gRPC handler which serves the
// "secured_service" service "also_doubly_secure" endpoint.
func NewAlsoDoublySecureHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeAlsoDoublySecureRequest, EncodeAlsoDoublySecureResponse)
	}
	return h
}

// AlsoDoublySecure implements the "AlsoDoublySecure" method in
// secured_servicepb.SecuredServiceServer interface.
func (s *Server) AlsoDoublySecure(ctx context.Context, message *secured_servicepb.AlsoDoublySecureRequest) (*secured_servicepb.AlsoDoublySecureResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "also_doubly_secure")
	ctx = context.WithValue(ctx, goa.ServiceKey, "secured_service")
	resp, err := s.AlsoDoublySecureH.Handle(ctx, message)
	if err != nil {
		var en goa.GoaErrorNamer
		if errors.As(err, &en) {
			switch en.GoaErrorName() {
			case "invalid-scopes":
				return nil, goagrpc.NewStatusError(codes.Unauthenticated, err, goagrpc.NewErrorResponse(err))
			case "unauthorized":
				return nil, goagrpc.NewStatusError(codes.Unauthenticated, err, goagrpc.NewErrorResponse(err))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*secured_servicepb.AlsoDoublySecureResponse), nil
}
