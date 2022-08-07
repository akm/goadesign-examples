// Code generated by goa v3.8.1, DO NOT EDIT.
//
// calc gRPC server encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/basic/design -o basic

package server

import (
	"context"

	calc "goa.design/examples/basic/gen/calc"
	calcpb "goa.design/examples/basic/gen/grpc/calc/pb"
	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc/metadata"
)

// EncodeMultiplyResponse encodes responses from the "calc" service "multiply"
// endpoint.
func EncodeMultiplyResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(int)
	if !ok {
		return nil, goagrpc.ErrInvalidType("calc", "multiply", "int", v)
	}
	resp := NewProtoMultiplyResponse(result)
	return resp, nil
}

// DecodeMultiplyRequest decodes requests sent to "calc" service "multiply"
// endpoint.
func DecodeMultiplyRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *calcpb.MultiplyRequest
		ok      bool
	)
	{
		if message, ok = v.(*calcpb.MultiplyRequest); !ok {
			return nil, goagrpc.ErrInvalidType("calc", "multiply", "*calcpb.MultiplyRequest", v)
		}
	}
	var payload *calc.MultiplyPayload
	{
		payload = NewMultiplyPayload(message)
	}
	return payload, nil
}
