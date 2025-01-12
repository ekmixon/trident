// Code generated by go-swagger; DO NOT EDIT.

package name_services

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

// NewNameMappingPositionGetParams creates a new NameMappingPositionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNameMappingPositionGetParams() *NameMappingPositionGetParams {
	return &NameMappingPositionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNameMappingPositionGetParamsWithTimeout creates a new NameMappingPositionGetParams object
// with the ability to set a timeout on a request.
func NewNameMappingPositionGetParamsWithTimeout(timeout time.Duration) *NameMappingPositionGetParams {
	return &NameMappingPositionGetParams{
		timeout: timeout,
	}
}

// NewNameMappingPositionGetParamsWithContext creates a new NameMappingPositionGetParams object
// with the ability to set a context for a request.
func NewNameMappingPositionGetParamsWithContext(ctx context.Context) *NameMappingPositionGetParams {
	return &NameMappingPositionGetParams{
		Context: ctx,
	}
}

// NewNameMappingPositionGetParamsWithHTTPClient creates a new NameMappingPositionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewNameMappingPositionGetParamsWithHTTPClient(client *http.Client) *NameMappingPositionGetParams {
	return &NameMappingPositionGetParams{
		HTTPClient: client,
	}
}

/* NameMappingPositionGetParams contains all the parameters to send to the API endpoint
   for the name mapping position get operation.

   Typically these are written to a http.Request.
*/
type NameMappingPositionGetParams struct {

	/* Direction.

	   Direction
	*/
	DirectionPathParameter string

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* Index.

	   Position of the entry in the list
	*/
	IndexPathParameter int64

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SVMUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the name mapping position get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NameMappingPositionGetParams) WithDefaults() *NameMappingPositionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the name mapping position get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NameMappingPositionGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the name mapping position get params
func (o *NameMappingPositionGetParams) WithTimeout(timeout time.Duration) *NameMappingPositionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the name mapping position get params
func (o *NameMappingPositionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the name mapping position get params
func (o *NameMappingPositionGetParams) WithContext(ctx context.Context) *NameMappingPositionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the name mapping position get params
func (o *NameMappingPositionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the name mapping position get params
func (o *NameMappingPositionGetParams) WithHTTPClient(client *http.Client) *NameMappingPositionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the name mapping position get params
func (o *NameMappingPositionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDirectionPathParameter adds the direction to the name mapping position get params
func (o *NameMappingPositionGetParams) WithDirectionPathParameter(direction string) *NameMappingPositionGetParams {
	o.SetDirectionPathParameter(direction)
	return o
}

// SetDirectionPathParameter adds the direction to the name mapping position get params
func (o *NameMappingPositionGetParams) SetDirectionPathParameter(direction string) {
	o.DirectionPathParameter = direction
}

// WithFieldsQueryParameter adds the fields to the name mapping position get params
func (o *NameMappingPositionGetParams) WithFieldsQueryParameter(fields []string) *NameMappingPositionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the name mapping position get params
func (o *NameMappingPositionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithIndexPathParameter adds the index to the name mapping position get params
func (o *NameMappingPositionGetParams) WithIndexPathParameter(index int64) *NameMappingPositionGetParams {
	o.SetIndexPathParameter(index)
	return o
}

// SetIndexPathParameter adds the index to the name mapping position get params
func (o *NameMappingPositionGetParams) SetIndexPathParameter(index int64) {
	o.IndexPathParameter = index
}

// WithSVMUUIDPathParameter adds the svmUUID to the name mapping position get params
func (o *NameMappingPositionGetParams) WithSVMUUIDPathParameter(svmUUID string) *NameMappingPositionGetParams {
	o.SetSVMUUIDPathParameter(svmUUID)
	return o
}

// SetSVMUUIDPathParameter adds the svmUuid to the name mapping position get params
func (o *NameMappingPositionGetParams) SetSVMUUIDPathParameter(svmUUID string) {
	o.SVMUUIDPathParameter = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *NameMappingPositionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param direction
	if err := r.SetPathParam("direction", o.DirectionPathParameter); err != nil {
		return err
	}

	if o.FieldsQueryParameter != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	// path param index
	if err := r.SetPathParam("index", swag.FormatInt64(o.IndexPathParameter)); err != nil {
		return err
	}

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SVMUUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamNameMappingPositionGet binds the parameter fields
func (o *NameMappingPositionGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.FieldsQueryParameter

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}
