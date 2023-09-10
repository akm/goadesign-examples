// Code generated by goa v3.13.0, DO NOT EDIT.
//
// calc gRPC server types
//
// Command:
// $ goa gen goa.design/examples/error/design -o error

package server

import (
	calc "goa.design/examples/error/gen/calc"
	calcpb "goa.design/examples/error/gen/grpc/calc/pb"
)

// NewDividePayload builds the payload of the "divide" endpoint of the "calc"
// service from the gRPC request type.
func NewDividePayload(message *calcpb.DivideRequest) *calc.DividePayload {
	v := &calc.DividePayload{
		Dividend: int(message.Dividend),
		Divisor:  int(message.Divisor),
	}
	return v
}

// NewProtoDivideResponse builds the gRPC response type from the result of the
// "divide" endpoint of the "calc" service.
func NewProtoDivideResponse(result *calc.DivideResult) *calcpb.DivideResponse {
	message := &calcpb.DivideResponse{
		Quotient: int32(result.Quotient),
		Reminder: int32(result.Reminder),
	}
	return message
}

// NewDivideDivByZeroError builds the gRPC error response type from the error
// of the "divide" endpoint of the "calc" service.
func NewDivideDivByZeroError(er *calc.DivByZero) *calcpb.DivideDivByZeroError {
	message := &calcpb.DivideDivByZeroError{
		Message_: er.Message,
	}
	return message
}
