// Code generated by goa v3.7.14, DO NOT EDIT.
//
// calc HTTP server types
//
// Command:
// $ goa gen goa.design/examples/basic/design -o basic

package server

import (
	calc "goa.design/examples/basic/gen/calc"
)

// NewMultiplyPayload builds a calc service multiply endpoint payload.
func NewMultiplyPayload(a int, b int) *calc.MultiplyPayload {
	v := &calc.MultiplyPayload{}
	v.A = a
	v.B = b

	return v
}
