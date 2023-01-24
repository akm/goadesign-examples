// Code generated by goa v3.11.0, DO NOT EDIT.
//
// HTTP request path constructors for the concat service.
//
// Command:
// $ goa gen goa.design/examples/encodings/cbor/design -o encodings/cbor

package client

import (
	"fmt"
)

// ConcatConcatPath returns the URL path to the concat service concat HTTP endpoint.
func ConcatConcatPath(a string, b string) string {
	return fmt.Sprintf("/concat/%v/%v", a, b)
}
