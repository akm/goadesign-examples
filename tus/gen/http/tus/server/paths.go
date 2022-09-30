// Code generated by goa v3.9.0, DO NOT EDIT.
//
// HTTP request path constructors for the tus service.
//
// Command:
// $ goa gen goa.design/examples/tus/design -o tus

package server

import (
	"fmt"
)

// HeadTusPath returns the URL path to the tus service head HTTP endpoint.
func HeadTusPath(id string) string {
	return fmt.Sprintf("/upload/%v", id)
}

// PatchTusPath returns the URL path to the tus service patch HTTP endpoint.
func PatchTusPath(id string) string {
	return fmt.Sprintf("/upload/%v", id)
}

// OptionsTusPath returns the URL path to the tus service options HTTP endpoint.
func OptionsTusPath() string {
	return "/upload"
}

// PostTusPath returns the URL path to the tus service post HTTP endpoint.
func PostTusPath() string {
	return "/upload"
}

// DeleteTusPath returns the URL path to the tus service delete HTTP endpoint.
func DeleteTusPath(id string) string {
	return fmt.Sprintf("/upload/%v", id)
}
