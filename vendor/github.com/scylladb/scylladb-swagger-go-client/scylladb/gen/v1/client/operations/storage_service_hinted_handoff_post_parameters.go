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
	"github.com/go-openapi/swag"
)

// NewStorageServiceHintedHandoffPostParams creates a new StorageServiceHintedHandoffPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStorageServiceHintedHandoffPostParams() *StorageServiceHintedHandoffPostParams {
	return &StorageServiceHintedHandoffPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStorageServiceHintedHandoffPostParamsWithTimeout creates a new StorageServiceHintedHandoffPostParams object
// with the ability to set a timeout on a request.
func NewStorageServiceHintedHandoffPostParamsWithTimeout(timeout time.Duration) *StorageServiceHintedHandoffPostParams {
	return &StorageServiceHintedHandoffPostParams{
		timeout: timeout,
	}
}

// NewStorageServiceHintedHandoffPostParamsWithContext creates a new StorageServiceHintedHandoffPostParams object
// with the ability to set a context for a request.
func NewStorageServiceHintedHandoffPostParamsWithContext(ctx context.Context) *StorageServiceHintedHandoffPostParams {
	return &StorageServiceHintedHandoffPostParams{
		Context: ctx,
	}
}

// NewStorageServiceHintedHandoffPostParamsWithHTTPClient creates a new StorageServiceHintedHandoffPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewStorageServiceHintedHandoffPostParamsWithHTTPClient(client *http.Client) *StorageServiceHintedHandoffPostParams {
	return &StorageServiceHintedHandoffPostParams{
		HTTPClient: client,
	}
}

/*
StorageServiceHintedHandoffPostParams contains all the parameters to send to the API endpoint

	for the storage service hinted handoff post operation.

	Typically these are written to a http.Request.
*/
type StorageServiceHintedHandoffPostParams struct {

	/* Throttle.

	   throttle in kb

	   Format: int32
	*/
	Throttle int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the storage service hinted handoff post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StorageServiceHintedHandoffPostParams) WithDefaults() *StorageServiceHintedHandoffPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the storage service hinted handoff post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StorageServiceHintedHandoffPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the storage service hinted handoff post params
func (o *StorageServiceHintedHandoffPostParams) WithTimeout(timeout time.Duration) *StorageServiceHintedHandoffPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage service hinted handoff post params
func (o *StorageServiceHintedHandoffPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage service hinted handoff post params
func (o *StorageServiceHintedHandoffPostParams) WithContext(ctx context.Context) *StorageServiceHintedHandoffPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage service hinted handoff post params
func (o *StorageServiceHintedHandoffPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage service hinted handoff post params
func (o *StorageServiceHintedHandoffPostParams) WithHTTPClient(client *http.Client) *StorageServiceHintedHandoffPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage service hinted handoff post params
func (o *StorageServiceHintedHandoffPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithThrottle adds the throttle to the storage service hinted handoff post params
func (o *StorageServiceHintedHandoffPostParams) WithThrottle(throttle int32) *StorageServiceHintedHandoffPostParams {
	o.SetThrottle(throttle)
	return o
}

// SetThrottle adds the throttle to the storage service hinted handoff post params
func (o *StorageServiceHintedHandoffPostParams) SetThrottle(throttle int32) {
	o.Throttle = throttle
}

// WriteToRequest writes these params to a swagger request
func (o *StorageServiceHintedHandoffPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param throttle
	qrThrottle := o.Throttle
	qThrottle := swag.FormatInt32(qrThrottle)
	if qThrottle != "" {

		if err := r.SetQueryParam("throttle", qThrottle); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}