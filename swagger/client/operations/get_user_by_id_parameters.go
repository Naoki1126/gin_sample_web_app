// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetUserByIDParams creates a new GetUserByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUserByIDParams() *GetUserByIDParams {
	return &GetUserByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserByIDParamsWithTimeout creates a new GetUserByIDParams object
// with the ability to set a timeout on a request.
func NewGetUserByIDParamsWithTimeout(timeout time.Duration) *GetUserByIDParams {
	return &GetUserByIDParams{
		timeout: timeout,
	}
}

// NewGetUserByIDParamsWithContext creates a new GetUserByIDParams object
// with the ability to set a context for a request.
func NewGetUserByIDParamsWithContext(ctx context.Context) *GetUserByIDParams {
	return &GetUserByIDParams{
		Context: ctx,
	}
}

// NewGetUserByIDParamsWithHTTPClient creates a new GetUserByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUserByIDParamsWithHTTPClient(client *http.Client) *GetUserByIDParams {
	return &GetUserByIDParams{
		HTTPClient: client,
	}
}

/* GetUserByIDParams contains all the parameters to send to the API endpoint
   for the get user by Id operation.

   Typically these are written to a http.Request.
*/
type GetUserByIDParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get user by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserByIDParams) WithDefaults() *GetUserByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get user by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get user by Id params
func (o *GetUserByIDParams) WithTimeout(timeout time.Duration) *GetUserByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user by Id params
func (o *GetUserByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user by Id params
func (o *GetUserByIDParams) WithContext(ctx context.Context) *GetUserByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user by Id params
func (o *GetUserByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user by Id params
func (o *GetUserByIDParams) WithHTTPClient(client *http.Client) *GetUserByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user by Id params
func (o *GetUserByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get user by Id params
func (o *GetUserByIDParams) WithID(id string) *GetUserByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get user by Id params
func (o *GetUserByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
