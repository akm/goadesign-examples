// Code generated by goa v3.1.2, DO NOT EDIT.
//
// HTTP request path constructors for the storage service.
//
// Command:
// $ goa gen goa.design/examples/cellar/design -o
// $(GOPATH)/src/goa.design/examples/cellar

package server

import (
	"fmt"
)

// ListStoragePath returns the URL path to the storage service list HTTP endpoint.
func ListStoragePath() string {
	return "/storage/"
}

// ShowStoragePath returns the URL path to the storage service show HTTP endpoint.
func ShowStoragePath(id string) string {
	return fmt.Sprintf("/storage/%v", id)
}

// AddStoragePath returns the URL path to the storage service add HTTP endpoint.
func AddStoragePath() string {
	return "/storage/"
}

// RemoveStoragePath returns the URL path to the storage service remove HTTP endpoint.
func RemoveStoragePath(id string) string {
	return fmt.Sprintf("/storage/%v", id)
}

// RateStoragePath returns the URL path to the storage service rate HTTP endpoint.
func RateStoragePath() string {
	return "/storage/rate"
}

// MultiAddStoragePath returns the URL path to the storage service multi_add HTTP endpoint.
func MultiAddStoragePath() string {
	return "/storage/multi_add"
}

// MultiUpdateStoragePath returns the URL path to the storage service multi_update HTTP endpoint.
func MultiUpdateStoragePath() string {
	return "/storage/multi_update"
}
