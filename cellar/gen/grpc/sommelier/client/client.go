// Code generated by goa v3.11.1, DO NOT EDIT.
//
// sommelier gRPC client
//
// Command:
// $ goa gen goa.design/examples/cellar/design -o cellar

package client

import (
	"context"

	sommelierpb "goa.design/examples/cellar/gen/grpc/sommelier/pb"
	goagrpc "goa.design/goa/v3/grpc"
	goapb "goa.design/goa/v3/grpc/pb"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc"
)

// Client lists the service endpoint gRPC clients.
type Client struct {
	grpccli sommelierpb.SommelierClient
	opts    []grpc.CallOption
}

// NewClient instantiates gRPC client for all the sommelier service servers.
func NewClient(cc *grpc.ClientConn, opts ...grpc.CallOption) *Client {
	return &Client{
		grpccli: sommelierpb.NewSommelierClient(cc),
		opts:    opts,
	}
}

// Pick calls the "Pick" function in sommelierpb.SommelierClient interface.
func (c *Client) Pick() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildPickFunc(c.grpccli, c.opts...),
			EncodePickRequest,
			DecodePickResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			resp := goagrpc.DecodeError(err)
			switch message := resp.(type) {
			case *goapb.ErrorResponse:
				return nil, goagrpc.NewServiceError(message)
			default:
				return nil, goa.Fault(err.Error())
			}
		}
		return res, nil
	}
}
