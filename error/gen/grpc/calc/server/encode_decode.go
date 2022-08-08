// Code generated by goa v3.8.2, DO NOT EDIT.
//
// calc gRPC server encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/error/design -o error

package server

import (
	"context"

	calc "goa.design/examples/error/gen/calc"
	calcpb "goa.design/examples/error/gen/grpc/calc/pb"
	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc/metadata"
)

// EncodeDivideResponse encodes responses from the "calc" service "divide"
// endpoint.
func EncodeDivideResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(*calc.DivideResult)
	if !ok {
		return nil, goagrpc.ErrInvalidType("calc", "divide", "*calc.DivideResult", v)
	}
	resp := NewProtoDivideResponse(result)
	return resp, nil
}

// DecodeDivideRequest decodes requests sent to "calc" service "divide"
// endpoint.
func DecodeDivideRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *calcpb.DivideRequest
		ok      bool
	)
	{
		if message, ok = v.(*calcpb.DivideRequest); !ok {
			return nil, goagrpc.ErrInvalidType("calc", "divide", "*calcpb.DivideRequest", v)
		}
	}
	var payload *calc.DividePayload
	{
		payload = NewDividePayload(message)
	}
	return payload, nil
}
