// Code generated by go-swagger; DO NOT EDIT.

package config

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

// NewFindConfigMemtableFlushQueueSizeParams creates a new FindConfigMemtableFlushQueueSizeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindConfigMemtableFlushQueueSizeParams() *FindConfigMemtableFlushQueueSizeParams {
	return &FindConfigMemtableFlushQueueSizeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigMemtableFlushQueueSizeParamsWithTimeout creates a new FindConfigMemtableFlushQueueSizeParams object
// with the ability to set a timeout on a request.
func NewFindConfigMemtableFlushQueueSizeParamsWithTimeout(timeout time.Duration) *FindConfigMemtableFlushQueueSizeParams {
	return &FindConfigMemtableFlushQueueSizeParams{
		timeout: timeout,
	}
}

// NewFindConfigMemtableFlushQueueSizeParamsWithContext creates a new FindConfigMemtableFlushQueueSizeParams object
// with the ability to set a context for a request.
func NewFindConfigMemtableFlushQueueSizeParamsWithContext(ctx context.Context) *FindConfigMemtableFlushQueueSizeParams {
	return &FindConfigMemtableFlushQueueSizeParams{
		Context: ctx,
	}
}

// NewFindConfigMemtableFlushQueueSizeParamsWithHTTPClient creates a new FindConfigMemtableFlushQueueSizeParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindConfigMemtableFlushQueueSizeParamsWithHTTPClient(client *http.Client) *FindConfigMemtableFlushQueueSizeParams {
	return &FindConfigMemtableFlushQueueSizeParams{
		HTTPClient: client,
	}
}

/*
FindConfigMemtableFlushQueueSizeParams contains all the parameters to send to the API endpoint

	for the find config memtable flush queue size operation.

	Typically these are written to a http.Request.
*/
type FindConfigMemtableFlushQueueSizeParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find config memtable flush queue size params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindConfigMemtableFlushQueueSizeParams) WithDefaults() *FindConfigMemtableFlushQueueSizeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find config memtable flush queue size params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindConfigMemtableFlushQueueSizeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find config memtable flush queue size params
func (o *FindConfigMemtableFlushQueueSizeParams) WithTimeout(timeout time.Duration) *FindConfigMemtableFlushQueueSizeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config memtable flush queue size params
func (o *FindConfigMemtableFlushQueueSizeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config memtable flush queue size params
func (o *FindConfigMemtableFlushQueueSizeParams) WithContext(ctx context.Context) *FindConfigMemtableFlushQueueSizeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config memtable flush queue size params
func (o *FindConfigMemtableFlushQueueSizeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config memtable flush queue size params
func (o *FindConfigMemtableFlushQueueSizeParams) WithHTTPClient(client *http.Client) *FindConfigMemtableFlushQueueSizeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config memtable flush queue size params
func (o *FindConfigMemtableFlushQueueSizeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigMemtableFlushQueueSizeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}