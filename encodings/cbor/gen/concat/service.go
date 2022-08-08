// Code generated by goa v3.8.2, DO NOT EDIT.
//
// concat service
//
// Command:
// $ goa gen goa.design/examples/encodings/cbor/design -o encodings/cbor

package concat

import (
	"context"
)

// The concat service performs operations on strings.

// The service uses the CBOR binary serialization standard to encode responses.
// It supports reading requests encoded with CBOR, JSON, XML or GOB.

type Service interface {
	// Concat implements concat.
	Concat(context.Context, *ConcatPayload) (res string, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "concat"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"concat"}

// ConcatPayload is the payload type of the concat service concat method.
type ConcatPayload struct {
	// Left operand
	A string
	// Right operand
	B string
}
