// Code generated by goa v3.8.2, DO NOT EDIT.
//
// storage HTTP server types
//
// Command:
// $ goa gen goa.design/examples/cellar/design -o cellar

package server

import (
	"unicode/utf8"

	storage "goa.design/examples/cellar/gen/storage"
	storageviews "goa.design/examples/cellar/gen/storage/views"
	goa "goa.design/goa/v3/pkg"
)

// AddRequestBody is the type of the "storage" service "add" endpoint HTTP
// request body.
type AddRequestBody struct {
	// Name of bottle
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Winery that produces wine
	Winery *WineryRequestBody `form:"winery,omitempty" json:"winery,omitempty" xml:"winery,omitempty"`
	// Vintage of bottle
	Vintage *uint32 `form:"vintage,omitempty" json:"vintage,omitempty" xml:"vintage,omitempty"`
	// Composition is the list of grape varietals and associated percentage.
	Composition []*ComponentRequestBody `form:"composition,omitempty" json:"composition,omitempty" xml:"composition,omitempty"`
	// Description of bottle
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Rating of bottle from 1 (worst) to 5 (best)
	Rating *uint32 `form:"rating,omitempty" json:"rating,omitempty" xml:"rating,omitempty"`
}

// MultiUpdateRequestBody is the type of the "storage" service "multi_update"
// endpoint HTTP request body.
type MultiUpdateRequestBody struct {
	// Array of bottle info that matches the ids attribute
	Bottles []*BottleRequestBody `form:"bottles,omitempty" json:"bottles,omitempty" xml:"bottles,omitempty"`
}

// StoredBottleResponseTinyCollection is the type of the "storage" service
// "list" endpoint HTTP response body.
type StoredBottleResponseTinyCollection []*StoredBottleResponseTiny

// ShowResponseBody is the type of the "storage" service "show" endpoint HTTP
// response body.
type ShowResponseBody struct {
	// ID is the unique id of the bottle.
	ID string `form:"id" json:"id" xml:"id"`
	// Name of bottle
	Name string `form:"name" json:"name" xml:"name"`
	// Winery that produces wine
	Winery *WineryResponseBodyTiny `form:"winery" json:"winery" xml:"winery"`
	// Vintage of bottle
	Vintage uint32 `form:"vintage" json:"vintage" xml:"vintage"`
	// Composition is the list of grape varietals and associated percentage.
	Composition []*ComponentResponseBody `form:"composition,omitempty" json:"composition,omitempty" xml:"composition,omitempty"`
	// Description of bottle
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Rating of bottle from 1 (worst) to 5 (best)
	Rating *uint32 `form:"rating,omitempty" json:"rating,omitempty" xml:"rating,omitempty"`
}

// ShowResponseBodyTiny is the type of the "storage" service "show" endpoint
// HTTP response body.
type ShowResponseBodyTiny struct {
	// ID is the unique id of the bottle.
	ID string `form:"id" json:"id" xml:"id"`
	// Name of bottle
	Name string `form:"name" json:"name" xml:"name"`
	// Winery that produces wine
	Winery *WineryResponseBodyTiny `form:"winery" json:"winery" xml:"winery"`
}

// ShowNotFoundResponseBody is the type of the "storage" service "show"
// endpoint HTTP response body for the "not_found" error.
type ShowNotFoundResponseBody struct {
	// Message of error
	Message string `form:"message" json:"message" xml:"message"`
	// ID of missing bottle
	ID string `form:"id" json:"id" xml:"id"`
}

// StoredBottleResponseTiny is used to define fields on response body types.
type StoredBottleResponseTiny struct {
	// ID is the unique id of the bottle.
	ID string `form:"id" json:"id" xml:"id"`
	// Name of bottle
	Name string `form:"name" json:"name" xml:"name"`
	// Winery that produces wine
	Winery *WineryResponseTiny `form:"winery" json:"winery" xml:"winery"`
}

// WineryResponseTiny is used to define fields on response body types.
type WineryResponseTiny struct {
	// Name of winery
	Name string `form:"name" json:"name" xml:"name"`
}

// WineryResponseBodyTiny is used to define fields on response body types.
type WineryResponseBodyTiny struct {
	// Name of winery
	Name string `form:"name" json:"name" xml:"name"`
}

// ComponentResponseBody is used to define fields on response body types.
type ComponentResponseBody struct {
	// Grape varietal
	Varietal string `form:"varietal" json:"varietal" xml:"varietal"`
	// Percentage of varietal in wine
	Percentage *uint32 `form:"percentage,omitempty" json:"percentage,omitempty" xml:"percentage,omitempty"`
}

// WineryRequestBody is used to define fields on request body types.
type WineryRequestBody struct {
	// Name of winery
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Region of winery
	Region *string `form:"region,omitempty" json:"region,omitempty" xml:"region,omitempty"`
	// Country of winery
	Country *string `form:"country,omitempty" json:"country,omitempty" xml:"country,omitempty"`
	// Winery website URL
	URL *string `form:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
}

// ComponentRequestBody is used to define fields on request body types.
type ComponentRequestBody struct {
	// Grape varietal
	Varietal *string `form:"varietal,omitempty" json:"varietal,omitempty" xml:"varietal,omitempty"`
	// Percentage of varietal in wine
	Percentage *uint32 `form:"percentage,omitempty" json:"percentage,omitempty" xml:"percentage,omitempty"`
}

// BottleRequestBody is used to define fields on request body types.
type BottleRequestBody struct {
	// Name of bottle
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Winery that produces wine
	Winery *WineryRequestBody `form:"winery,omitempty" json:"winery,omitempty" xml:"winery,omitempty"`
	// Vintage of bottle
	Vintage *uint32 `form:"vintage,omitempty" json:"vintage,omitempty" xml:"vintage,omitempty"`
	// Composition is the list of grape varietals and associated percentage.
	Composition []*ComponentRequestBody `form:"composition,omitempty" json:"composition,omitempty" xml:"composition,omitempty"`
	// Description of bottle
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Rating of bottle from 1 (worst) to 5 (best)
	Rating *uint32 `form:"rating,omitempty" json:"rating,omitempty" xml:"rating,omitempty"`
}

// NewStoredBottleResponseTinyCollection builds the HTTP response body from the
// result of the "list" endpoint of the "storage" service.
func NewStoredBottleResponseTinyCollection(res storageviews.StoredBottleCollectionView) StoredBottleResponseTinyCollection {
	body := make([]*StoredBottleResponseTiny, len(res))
	for i, val := range res {
		body[i] = marshalStorageviewsStoredBottleViewToStoredBottleResponseTiny(val)
	}
	return body
}

// NewShowResponseBody builds the HTTP response body from the result of the
// "show" endpoint of the "storage" service.
func NewShowResponseBody(res *storageviews.StoredBottleView) *ShowResponseBody {
	body := &ShowResponseBody{
		ID:          *res.ID,
		Name:        *res.Name,
		Vintage:     *res.Vintage,
		Description: res.Description,
		Rating:      res.Rating,
	}
	if res.Winery != nil {
		body.Winery = marshalStorageviewsWineryViewToWineryResponseBodyTiny(res.Winery)
	}
	if res.Composition != nil {
		body.Composition = make([]*ComponentResponseBody, len(res.Composition))
		for i, val := range res.Composition {
			body.Composition[i] = marshalStorageviewsComponentViewToComponentResponseBody(val)
		}
	}
	return body
}

// NewShowResponseBodyTiny builds the HTTP response body from the result of the
// "show" endpoint of the "storage" service.
func NewShowResponseBodyTiny(res *storageviews.StoredBottleView) *ShowResponseBodyTiny {
	body := &ShowResponseBodyTiny{
		ID:   *res.ID,
		Name: *res.Name,
	}
	if res.Winery != nil {
		body.Winery = marshalStorageviewsWineryViewToWineryResponseBodyTiny(res.Winery)
	}
	return body
}

// NewShowNotFoundResponseBody builds the HTTP response body from the result of
// the "show" endpoint of the "storage" service.
func NewShowNotFoundResponseBody(res *storage.NotFound) *ShowNotFoundResponseBody {
	body := &ShowNotFoundResponseBody{
		Message: res.Message,
		ID:      res.ID,
	}
	return body
}

// NewShowPayload builds a storage service show endpoint payload.
func NewShowPayload(id string, view *string) *storage.ShowPayload {
	v := &storage.ShowPayload{}
	v.ID = id
	v.View = view

	return v
}

// NewAddBottle builds a storage service add endpoint payload.
func NewAddBottle(body *AddRequestBody) *storage.Bottle {
	v := &storage.Bottle{
		Name:        *body.Name,
		Vintage:     *body.Vintage,
		Description: body.Description,
		Rating:      body.Rating,
	}
	v.Winery = unmarshalWineryRequestBodyToStorageWinery(body.Winery)
	if body.Composition != nil {
		v.Composition = make([]*storage.Component, len(body.Composition))
		for i, val := range body.Composition {
			v.Composition[i] = unmarshalComponentRequestBodyToStorageComponent(val)
		}
	}

	return v
}

// NewRemovePayload builds a storage service remove endpoint payload.
func NewRemovePayload(id string) *storage.RemovePayload {
	v := &storage.RemovePayload{}
	v.ID = id

	return v
}

// NewMultiAddBottle builds a storage service multi_add endpoint payload.
func NewMultiAddBottle(body []*BottleRequestBody) []*storage.Bottle {
	v := make([]*storage.Bottle, len(body))
	for i, val := range body {
		v[i] = unmarshalBottleRequestBodyToStorageBottle(val)
	}
	return v
}

// NewMultiUpdatePayload builds a storage service multi_update endpoint payload.
func NewMultiUpdatePayload(body *MultiUpdateRequestBody, ids []string) *storage.MultiUpdatePayload {
	v := &storage.MultiUpdatePayload{}
	v.Bottles = make([]*storage.Bottle, len(body.Bottles))
	for i, val := range body.Bottles {
		v.Bottles[i] = unmarshalBottleRequestBodyToStorageBottle(val)
	}
	v.Ids = ids

	return v
}

// ValidateAddRequestBody runs the validations defined on AddRequestBody
func ValidateAddRequestBody(body *AddRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Winery == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("winery", "body"))
	}
	if body.Vintage == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("vintage", "body"))
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 100, false))
		}
	}
	if body.Winery != nil {
		if err2 := ValidateWineryRequestBody(body.Winery); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.Vintage != nil {
		if *body.Vintage < 1900 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.vintage", *body.Vintage, 1900, true))
		}
	}
	if body.Vintage != nil {
		if *body.Vintage > 2020 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.vintage", *body.Vintage, 2020, false))
		}
	}
	for _, e := range body.Composition {
		if e != nil {
			if err2 := ValidateComponentRequestBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if body.Description != nil {
		if utf8.RuneCountInString(*body.Description) > 2000 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.description", *body.Description, utf8.RuneCountInString(*body.Description), 2000, false))
		}
	}
	if body.Rating != nil {
		if *body.Rating < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.rating", *body.Rating, 1, true))
		}
	}
	if body.Rating != nil {
		if *body.Rating > 5 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.rating", *body.Rating, 5, false))
		}
	}
	return
}

// ValidateMultiUpdateRequestBody runs the validations defined on
// multi_update_request_body
func ValidateMultiUpdateRequestBody(body *MultiUpdateRequestBody) (err error) {
	if body.Bottles == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("bottles", "body"))
	}
	for _, e := range body.Bottles {
		if e != nil {
			if err2 := ValidateBottleRequestBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateWineryRequestBody runs the validations defined on WineryRequestBody
func ValidateWineryRequestBody(body *WineryRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Region == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("region", "body"))
	}
	if body.Country == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("country", "body"))
	}
	if body.Region != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.region", *body.Region, "[a-zA-Z '\\.]+"))
	}
	if body.Country != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.country", *body.Country, "[a-zA-Z '\\.]+"))
	}
	if body.URL != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.url", *body.URL, "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$"))
	}
	return
}

// ValidateComponentRequestBody runs the validations defined on
// ComponentRequestBody
func ValidateComponentRequestBody(body *ComponentRequestBody) (err error) {
	if body.Varietal == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("varietal", "body"))
	}
	if body.Varietal != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.varietal", *body.Varietal, "[A-Za-z' ]+"))
	}
	if body.Varietal != nil {
		if utf8.RuneCountInString(*body.Varietal) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.varietal", *body.Varietal, utf8.RuneCountInString(*body.Varietal), 100, false))
		}
	}
	if body.Percentage != nil {
		if *body.Percentage < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.percentage", *body.Percentage, 1, true))
		}
	}
	if body.Percentage != nil {
		if *body.Percentage > 100 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.percentage", *body.Percentage, 100, false))
		}
	}
	return
}

// ValidateBottleRequestBody runs the validations defined on BottleRequestBody
func ValidateBottleRequestBody(body *BottleRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Winery == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("winery", "body"))
	}
	if body.Vintage == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("vintage", "body"))
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 100, false))
		}
	}
	if body.Winery != nil {
		if err2 := ValidateWineryRequestBody(body.Winery); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.Vintage != nil {
		if *body.Vintage < 1900 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.vintage", *body.Vintage, 1900, true))
		}
	}
	if body.Vintage != nil {
		if *body.Vintage > 2020 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.vintage", *body.Vintage, 2020, false))
		}
	}
	for _, e := range body.Composition {
		if e != nil {
			if err2 := ValidateComponentRequestBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if body.Description != nil {
		if utf8.RuneCountInString(*body.Description) > 2000 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.description", *body.Description, utf8.RuneCountInString(*body.Description), 2000, false))
		}
	}
	if body.Rating != nil {
		if *body.Rating < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.rating", *body.Rating, 1, true))
		}
	}
	if body.Rating != nil {
		if *body.Rating > 5 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.rating", *body.Rating, 5, false))
		}
	}
	return
}
