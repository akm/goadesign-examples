// Code generated by goa v3.5.2, DO NOT EDIT.
//
// calc gRPC client
//
// Command:
// $ goa gen goa.design/examples/error/design -o
// $(GOPATH)/src/goa.design/examples/error

package client

import (
	"context"

	calcpb "goa.design/examples/error/gen/grpc/calc/pb"
	goagrpc "goa.design/goa/v3/grpc"
	goapb "goa.design/goa/v3/grpc/pb"
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

// Divide calls the "Divide" function in calcpb.CalcClient interface.
func (c *Client) Divide() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildDivideFunc(c.grpccli, c.opts...),
			EncodeDivideRequest,
			DecodeDivideResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			resp := goagrpc.DecodeError(err)
			switch message := resp.(type) {
			case *calcpb.DivideDivByZeroError:
				return nil, NewDivideDivByZeroError(message)
			case *goapb.ErrorResponse:
				return nil, goagrpc.NewServiceError(message)
			default:
				return nil, goa.Fault(err.Error())
			}
		}
		return res, nil
	}
}
