// Code generated by goa v3.11.1, DO NOT EDIT.
//
// sommelier gRPC server types
//
// Command:
// $ goa gen goa.design/examples/cellar/design -o cellar

package server

import (
	sommelierpb "goa.design/examples/cellar/gen/grpc/sommelier/pb"
	sommelier "goa.design/examples/cellar/gen/sommelier"
	sommelierviews "goa.design/examples/cellar/gen/sommelier/views"
)

// NewPickPayload builds the payload of the "pick" endpoint of the "sommelier"
// service from the gRPC request type.
func NewPickPayload(message *sommelierpb.PickRequest) *sommelier.Criteria {
	v := &sommelier.Criteria{
		Name:   message.Name,
		Winery: message.Winery,
	}
	if message.Varietal != nil {
		v.Varietal = make([]string, len(message.Varietal))
		for i, val := range message.Varietal {
			v.Varietal[i] = val
		}
	}
	return v
}

// NewProtoStoredBottleCollection builds the gRPC response type from the result
// of the "pick" endpoint of the "sommelier" service.
func NewProtoStoredBottleCollection(result sommelierviews.StoredBottleCollectionView) *sommelierpb.StoredBottleCollection {
	message := &sommelierpb.StoredBottleCollection{}
	message.Field = make([]*sommelierpb.StoredBottle, len(result))
	for i, val := range result {
		message.Field[i] = &sommelierpb.StoredBottle{
			Id:          *val.ID,
			Name:        *val.Name,
			Vintage:     *val.Vintage,
			Description: val.Description,
			Rating:      val.Rating,
		}
		if val.Winery != nil {
			message.Field[i].Winery = svcSommelierviewsWineryViewToSommelierpbWinery(val.Winery)
		}
		if val.Composition != nil {
			message.Field[i].Composition = make([]*sommelierpb.Component, len(val.Composition))
			for j, val := range val.Composition {
				message.Field[i].Composition[j] = &sommelierpb.Component{
					Varietal:   *val.Varietal,
					Percentage: val.Percentage,
				}
			}
		}
	}
	return message
}

// svcSommelierviewsWineryViewToSommelierpbWinery builds a value of type
// *sommelierpb.Winery from a value of type *sommelierviews.WineryView.
func svcSommelierviewsWineryViewToSommelierpbWinery(v *sommelierviews.WineryView) *sommelierpb.Winery {
	res := &sommelierpb.Winery{
		Name:    *v.Name,
		Region:  *v.Region,
		Country: *v.Country,
		Url:     v.URL,
	}

	return res
}

// protobufSommelierpbWineryToSommelierviewsWineryView builds a value of type
// *sommelierviews.WineryView from a value of type *sommelierpb.Winery.
func protobufSommelierpbWineryToSommelierviewsWineryView(v *sommelierpb.Winery) *sommelierviews.WineryView {
	res := &sommelierviews.WineryView{
		Name:    &v.Name,
		Region:  &v.Region,
		Country: &v.Country,
		URL:     v.Url,
	}

	return res
}
