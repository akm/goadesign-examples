// Code generated by goa v3.10.2, DO NOT EDIT.
//
// sommelier views
//
// Command:
// $ goa gen goa.design/examples/cellar/design -o cellar

package views

import (
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// StoredBottleCollection is the viewed result type that is projected based on
// a view.
type StoredBottleCollection struct {
	// Type to project
	Projected StoredBottleCollectionView
	// View to render
	View string
}

// StoredBottleCollectionView is a type that runs validations on a projected
// type.
type StoredBottleCollectionView []*StoredBottleView

// StoredBottleView is a type that runs validations on a projected type.
type StoredBottleView struct {
	// ID is the unique id of the bottle.
	ID *string
	// Name of bottle
	Name *string
	// Winery that produces wine
	Winery *WineryView
	// Vintage of bottle
	Vintage *uint32
	// Composition is the list of grape varietals and associated percentage.
	Composition []*ComponentView
	// Description of bottle
	Description *string
	// Rating of bottle from 1 (worst) to 5 (best)
	Rating *uint32
}

// WineryView is a type that runs validations on a projected type.
type WineryView struct {
	// Name of winery
	Name *string
	// Region of winery
	Region *string
	// Country of winery
	Country *string
	// Winery website URL
	URL *string
}

// ComponentView is a type that runs validations on a projected type.
type ComponentView struct {
	// Grape varietal
	Varietal *string
	// Percentage of varietal in wine
	Percentage *uint32
}

var (
	// StoredBottleCollectionMap is a map indexing the attribute names of
	// StoredBottleCollection by view name.
	StoredBottleCollectionMap = map[string][]string{
		"default": {
			"id",
			"name",
			"winery",
			"vintage",
			"composition",
			"description",
			"rating",
		},
		"tiny": {
			"id",
			"name",
			"winery",
		},
	}
	// StoredBottleMap is a map indexing the attribute names of StoredBottle by
	// view name.
	StoredBottleMap = map[string][]string{
		"default": {
			"id",
			"name",
			"winery",
			"vintage",
			"composition",
			"description",
			"rating",
		},
		"tiny": {
			"id",
			"name",
			"winery",
		},
	}
	// WineryMap is a map indexing the attribute names of Winery by view name.
	WineryMap = map[string][]string{
		"default": {
			"name",
			"region",
			"country",
			"url",
		},
		"tiny": {
			"name",
		},
	}
)

// ValidateStoredBottleCollection runs the validations defined on the viewed
// result type StoredBottleCollection.
func ValidateStoredBottleCollection(result StoredBottleCollection) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateStoredBottleCollectionView(result.Projected)
	case "tiny":
		err = ValidateStoredBottleCollectionViewTiny(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default", "tiny"})
	}
	return
}

// ValidateStoredBottleCollectionView runs the validations defined on
// StoredBottleCollectionView using the "default" view.
func ValidateStoredBottleCollectionView(result StoredBottleCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateStoredBottleView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateStoredBottleCollectionViewTiny runs the validations defined on
// StoredBottleCollectionView using the "tiny" view.
func ValidateStoredBottleCollectionViewTiny(result StoredBottleCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateStoredBottleViewTiny(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateStoredBottleView runs the validations defined on StoredBottleView
// using the "default" view.
func ValidateStoredBottleView(result *StoredBottleView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Vintage == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("vintage", "result"))
	}
	if result.Name != nil {
		if utf8.RuneCountInString(*result.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.name", *result.Name, utf8.RuneCountInString(*result.Name), 100, false))
		}
	}
	if result.Vintage != nil {
		if *result.Vintage < 1900 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("result.vintage", *result.Vintage, 1900, true))
		}
	}
	if result.Vintage != nil {
		if *result.Vintage > 2020 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("result.vintage", *result.Vintage, 2020, false))
		}
	}
	for _, e := range result.Composition {
		if e != nil {
			if err2 := ValidateComponentView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if result.Description != nil {
		if utf8.RuneCountInString(*result.Description) > 2000 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.description", *result.Description, utf8.RuneCountInString(*result.Description), 2000, false))
		}
	}
	if result.Rating != nil {
		if *result.Rating < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("result.rating", *result.Rating, 1, true))
		}
	}
	if result.Rating != nil {
		if *result.Rating > 5 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("result.rating", *result.Rating, 5, false))
		}
	}
	if result.Winery != nil {
		if err2 := ValidateWineryViewTiny(result.Winery); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateStoredBottleViewTiny runs the validations defined on
// StoredBottleView using the "tiny" view.
func ValidateStoredBottleViewTiny(result *StoredBottleView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Name != nil {
		if utf8.RuneCountInString(*result.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.name", *result.Name, utf8.RuneCountInString(*result.Name), 100, false))
		}
	}
	if result.Winery != nil {
		if err2 := ValidateWineryViewTiny(result.Winery); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateWineryView runs the validations defined on WineryView using the
// "default" view.
func ValidateWineryView(result *WineryView) (err error) {
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Region == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("region", "result"))
	}
	if result.Country == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("country", "result"))
	}
	if result.Region != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("result.region", *result.Region, "[a-zA-Z '\\.]+"))
	}
	if result.Country != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("result.country", *result.Country, "[a-zA-Z '\\.]+"))
	}
	if result.URL != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("result.url", *result.URL, "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$"))
	}
	return
}

// ValidateWineryViewTiny runs the validations defined on WineryView using the
// "tiny" view.
func ValidateWineryViewTiny(result *WineryView) (err error) {
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	return
}

// ValidateComponentView runs the validations defined on ComponentView.
func ValidateComponentView(result *ComponentView) (err error) {
	if result.Varietal == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("varietal", "result"))
	}
	if result.Varietal != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("result.varietal", *result.Varietal, "[A-Za-z' ]+"))
	}
	if result.Varietal != nil {
		if utf8.RuneCountInString(*result.Varietal) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.varietal", *result.Varietal, utf8.RuneCountInString(*result.Varietal), 100, false))
		}
	}
	if result.Percentage != nil {
		if *result.Percentage < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("result.percentage", *result.Percentage, 1, true))
		}
	}
	if result.Percentage != nil {
		if *result.Percentage > 100 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("result.percentage", *result.Percentage, 100, false))
		}
	}
	return
}
