// Code generated by goa v3.8.2, DO NOT EDIT.
//
// tus service
//
// Command:
// $ goa gen goa.design/examples/tus/design -o tus

package tus

import (
	"context"
	"io"

	goa "goa.design/goa/v3/pkg"
)

// The tus service exposes the methods required to implement the tus protocol
type Service interface {
	// Clients use the HEAD request to determine the offset at which the upload
	// should be continued.
	Head(context.Context, *HeadPayload) (res *HeadResult, err error)
	// Clients use the PATCH method to start or resume an upload.
	Patch(context.Context, *PatchPayload, io.ReadCloser) (res *PatchResult, err error)
	// Clients use the OPTIONS method to gather information about the Server’s
	// current configuration.
	Options(context.Context) (res *OptionsResult, err error)
	// Clients use the POST method against a known upload creation URL to request a
	// new upload resource.
	Post(context.Context, *PostPayload, io.ReadCloser) (res *PostResult, err error)
	// Clients use the DELETE method to terminate completed and unfinished uploads
	// allowing the Server to free up used resources.
	Delete(context.Context, *DeletePayload) (res *DeleteResult, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "tus"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [5]string{"head", "patch", "options", "post", "delete"}

// DeletePayload is the payload type of the tus service delete method.
type DeletePayload struct {
	// IDs are generated using Xid: https://github.com/rs/xid
	ID string
	// tusResumable represents a tus protocol version.
	TusResumable string
}

// DeleteResult is the result type of the tus service delete method.
type DeleteResult struct {
	// tusResumable represents a tus protocol version.
	TusResumable string
}

// ErrInvalidTUSResumable describes the error returned when a non-supported
// Tus-Resumable header is provided by the client.
type ErrInvalidTUSResumable struct {
	// Comma separated list of supported versions.
	TusVersion string
}

// HeadPayload is the payload type of the tus service head method.
type HeadPayload struct {
	// IDs are generated using Xid: https://github.com/rs/xid
	ID string
	// tusResumable represents a tus protocol version.
	TusResumable string
}

// HeadResult is the result type of the tus service head method.
type HeadResult struct {
	// tusResumable represents a tus protocol version.
	TusResumable string
	// uploadOffset represents a byte offset within a resource.
	UploadOffset int64
	// uploadLength represents the size of the entire upload in bytes.
	UploadLength *int64
	// The Upload-Defer-Length request and response header indicates that the size
	// of the upload is not known currently and will be transferred later.
	UploadDeferLength *int
	// The Client MAY supply the Upload-Metadata header to add additional metadata
	// to the upload creation request.
	UploadMetadata *string
}

// OptionsResult is the result type of the tus service options method.
type OptionsResult struct {
	// tusResumable represents a tus protocol version.
	TusResumable string
	// tusVersion is a comma separated list of protocol versions supported by the
	// server. This implementation only supports 1.0.0.
	TusVersion string
	// tusExtension is a comma separated list of extensions supported by the
	// server. This implementation supports the creation, creation-with-upload,
	// expiration, checksum and termination extensions
	TusExtension string
	// tusMaxSize represents the maximum allowed size of an entire upload in bytes.
	TusMaxSize *int64
	// A Client MAY include the Upload-Checksum header in a PATCH request. Once the
	// entire request has been received, the Server MUST verify the uploaded chunk
	// against the provided checksum using the specified algorithm.
	TusChecksumAlgorithm string
}

// PatchPayload is the payload type of the tus service patch method.
type PatchPayload struct {
	// IDs are generated using Xid: https://github.com/rs/xid
	ID string
	// tusResumable represents a tus protocol version.
	TusResumable string
	// uploadOffset represents a byte offset within a resource.
	UploadOffset int64
	// A Client MAY include the Upload-Checksum header in a PATCH request. Once the
	// entire request has been received, the Server MUST verify the uploaded chunk
	// against the provided checksum using the specified algorithm.
	UploadChecksum *string
}

// PatchResult is the result type of the tus service patch method.
type PatchResult struct {
	// tusResumable represents a tus protocol version.
	TusResumable string
	// uploadOffset represents a byte offset within a resource.
	UploadOffset int64
	// The Upload-Expires response header indicates the time after which the
	// unfinished upload expires.
	UploadExpires *string
}

// PostPayload is the payload type of the tus service post method.
type PostPayload struct {
	// tusResumable represents a tus protocol version.
	TusResumable string
	// uploadLength represents the size of the entire upload in bytes.
	UploadLength *int64
	// The Upload-Defer-Length request and response header indicates that the size
	// of the upload is not known currently and will be transferred later.
	UploadDeferLength *int
	// A Client MAY include the Upload-Checksum header in a PATCH request. Once the
	// entire request has been received, the Server MUST verify the uploaded chunk
	// against the provided checksum using the specified algorithm.
	UploadChecksum *string
	// The Client MAY supply the Upload-Metadata header to add additional metadata
	// to the upload creation request.
	UploadMetadata *string
	// Length of the upload
	TusMaxSize *int64
}

// PostResult is the result type of the tus service post method.
type PostResult struct {
	// tusResumable represents a tus protocol version.
	TusResumable string
	// uploadOffset represents a byte offset within a resource.
	UploadOffset int64
	// The Upload-Expires response header indicates the time after which the
	// unfinished upload expires.
	UploadExpires *string
	// URL of created resource
	Location string
}

// Error returns an error description.
func (e *ErrInvalidTUSResumable) Error() string {
	return "ErrInvalidTUSResumable describes the error returned when a non-supported Tus-Resumable header is provided by the client."
}

// ErrorName returns "ErrInvalidTUSResumable".
func (e *ErrInvalidTUSResumable) ErrorName() string {
	return "InvalidTusResumable"
}

// MakeNotFound builds a goa.ServiceError from an error.
func MakeNotFound(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "NotFound", false, false, false)
}

// MakeGone builds a goa.ServiceError from an error.
func MakeGone(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "Gone", false, false, false)
}

// MakeInvalidContentType builds a goa.ServiceError from an error.
func MakeInvalidContentType(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "InvalidContentType", false, false, false)
}

// MakeInvalidOffset builds a goa.ServiceError from an error.
func MakeInvalidOffset(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "InvalidOffset", false, false, false)
}

// MakeInvalidChecksumAlgorithm builds a goa.ServiceError from an error.
func MakeInvalidChecksumAlgorithm(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "InvalidChecksumAlgorithm", false, false, false)
}

// MakeChecksumMismatch builds a goa.ServiceError from an error.
func MakeChecksumMismatch(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "ChecksumMismatch", false, false, false)
}

// MakeInternal builds a goa.ServiceError from an error.
func MakeInternal(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "Internal", false, false, false)
}

// MakeMissingHeader builds a goa.ServiceError from an error.
func MakeMissingHeader(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "MissingHeader", false, false, false)
}

// MakeInvalidDeferLength builds a goa.ServiceError from an error.
func MakeInvalidDeferLength(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "InvalidDeferLength", false, false, false)
}

// MakeMaximumSizeExceeded builds a goa.ServiceError from an error.
func MakeMaximumSizeExceeded(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "MaximumSizeExceeded", false, false, false)
}
