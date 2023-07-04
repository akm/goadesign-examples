// Code generated by goa v3.12.1, DO NOT EDIT.
//
// calc gRPC client
//
// Command:
// $ goa gen goa.design/examples/basic/design -o basic

package client

import (
	"context"

	calcpb "goa.design/examples/basic/gen/grpc/calc/pb"
	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc"
)

// Client lists the service endpoint gRPC clients.
type Client struct {
	grpccli calcpb.CalcClient
	opts    []grpc.CallOption
}

// NewClient instantiates gRPC client for all the calc service servers.
func NewClient(cc *grpc.ClientConn, opts ...grpc.CallOption) *Client {
	return &Client{
		grpccli: calcpb.NewCalcClient(cc),
		opts:    opts,
	}
}

// Multiply calls the "Multiply" function in calcpb.CalcClient interface.
func (c *Client) Multiply() goa.Endpoint {
	return func(ctx context.Context, v any) (any, error) {
		inv := goagrpc.NewInvoker(
			BuildMultiplyFunc(c.grpccli, c.opts...),
			EncodeMultiplyRequest,
			DecodeMultiplyResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}
