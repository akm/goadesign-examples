// Code generated by goa v3.2.6, DO NOT EDIT.
//
// updown HTTP server types
//
// Command:
// $ goa gen goa.design/examples/upload_download/design -o
// $(GOPATH)/src/goa.design/examples/upload_download

package server

import (
	updown "goa.design/examples/upload_download/gen/updown"
	goa "goa.design/goa/v3/pkg"
)

// UploadInvalidMediaTypeResponseBody is the type of the "updown" service
// "upload" endpoint HTTP response body for the "invalid_media_type" error.
type UploadInvalidMediaTypeResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UploadInvalidMultipartRequestResponseBody is the type of the "updown"
// service "upload" endpoint HTTP response body for the
// "invalid_multipart_request" error.
type UploadInvalidMultipartRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UploadInternalErrorResponseBody is the type of the "updown" service "upload"
// endpoint HTTP response body for the "internal_error" error.
type UploadInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// DownloadInvalidFilePathResponseBody is the type of the "updown" service
// "download" endpoint HTTP response body for the "invalid_file_path" error.
type DownloadInvalidFilePathResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// DownloadInternalErrorResponseBody is the type of the "updown" service
// "download" endpoint HTTP response body for the "internal_error" error.
type DownloadInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// NewUploadInvalidMediaTypeResponseBody builds the HTTP response body from the
// result of the "upload" endpoint of the "updown" service.
func NewUploadInvalidMediaTypeResponseBody(res *goa.ServiceError) *UploadInvalidMediaTypeResponseBody {
	body := &UploadInvalidMediaTypeResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUploadInvalidMultipartRequestResponseBody builds the HTTP response body
// from the result of the "upload" endpoint of the "updown" service.
func NewUploadInvalidMultipartRequestResponseBody(res *goa.ServiceError) *UploadInvalidMultipartRequestResponseBody {
	body := &UploadInvalidMultipartRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUploadInternalErrorResponseBody builds the HTTP response body from the
// result of the "upload" endpoint of the "updown" service.
func NewUploadInternalErrorResponseBody(res *goa.ServiceError) *UploadInternalErrorResponseBody {
	body := &UploadInternalErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewDownloadInvalidFilePathResponseBody builds the HTTP response body from
// the result of the "download" endpoint of the "updown" service.
func NewDownloadInvalidFilePathResponseBody(res *goa.ServiceError) *DownloadInvalidFilePathResponseBody {
	body := &DownloadInvalidFilePathResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewDownloadInternalErrorResponseBody builds the HTTP response body from the
// result of the "download" endpoint of the "updown" service.
func NewDownloadInternalErrorResponseBody(res *goa.ServiceError) *DownloadInternalErrorResponseBody {
	body := &DownloadInternalErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUploadPayload builds a updown service upload endpoint payload.
func NewUploadPayload(dir string, contentType string) *updown.UploadPayload {
	v := &updown.UploadPayload{}
	v.Dir = dir
	v.ContentType = contentType

	return v
}
