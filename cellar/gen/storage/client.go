// Code generated by goa v3.11.0, DO NOT EDIT.
//
// storage client
//
// Command:
// $ goa gen goa.design/examples/cellar/design -o cellar

package storage

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "storage" service client.
type Client struct {
	ListEndpoint        goa.Endpoint
	ShowEndpoint        goa.Endpoint
	AddEndpoint         goa.Endpoint
	RemoveEndpoint      goa.Endpoint
	RateEndpoint        goa.Endpoint
	MultiAddEndpoint    goa.Endpoint
	MultiUpdateEndpoint goa.Endpoint
}

// NewClient initializes a "storage" service client given the endpoints.
func NewClient(list, show, add, remove, rate, multiAdd, multiUpdate goa.Endpoint) *Client {
	return &Client{
		ListEndpoint:        list,
		ShowEndpoint:        show,
		AddEndpoint:         add,
		RemoveEndpoint:      remove,
		RateEndpoint:        rate,
		MultiAddEndpoint:    multiAdd,
		MultiUpdateEndpoint: multiUpdate,
	}
}

// List calls the "list" endpoint of the "storage" service.
func (c *Client) List(ctx context.Context) (res StoredBottleCollection, err error) {
	var ires interface{}
	ires, err = c.ListEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(StoredBottleCollection), nil
}

// Show calls the "show" endpoint of the "storage" service.
// Show may return the following errors:
//   - "not_found" (type *NotFound): Bottle not found
//   - error: internal error
func (c *Client) Show(ctx context.Context, p *ShowPayload) (res *StoredBottle, err error) {
	var ires interface{}
	ires, err = c.ShowEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*StoredBottle), nil
}

// Add calls the "add" endpoint of the "storage" service.
func (c *Client) Add(ctx context.Context, p *Bottle) (res string, err error) {
	var ires interface{}
	ires, err = c.AddEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}

// Remove calls the "remove" endpoint of the "storage" service.
// Remove may return the following errors:
//   - "not_found" (type *NotFound): Bottle not found
//   - error: internal error
func (c *Client) Remove(ctx context.Context, p *RemovePayload) (err error) {
	_, err = c.RemoveEndpoint(ctx, p)
	return
}

// Rate calls the "rate" endpoint of the "storage" service.
func (c *Client) Rate(ctx context.Context, p map[uint32][]string) (err error) {
	_, err = c.RateEndpoint(ctx, p)
	return
}

// MultiAdd calls the "multi_add" endpoint of the "storage" service.
func (c *Client) MultiAdd(ctx context.Context, p []*Bottle) (res []string, err error) {
	var ires interface{}
	ires, err = c.MultiAddEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.([]string), nil
}

// MultiUpdate calls the "multi_update" endpoint of the "storage" service.
func (c *Client) MultiUpdate(ctx context.Context, p *MultiUpdatePayload) (err error) {
	_, err = c.MultiUpdateEndpoint(ctx, p)
	return
}
