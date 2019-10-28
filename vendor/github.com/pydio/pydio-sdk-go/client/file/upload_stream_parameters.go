// Code generated by go-swagger; DO NOT EDIT.

package file

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/pydio/pydio-sdk-go/models"
)

// NewUploadStreamParams creates a new UploadStreamParams object
// with the default values initialized.
func NewUploadStreamParams() *UploadStreamParams {
	var ()
	return &UploadStreamParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUploadStreamParamsWithTimeout creates a new UploadStreamParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUploadStreamParamsWithTimeout(timeout time.Duration) *UploadStreamParams {
	var ()
	return &UploadStreamParams{

		timeout: timeout,
	}
}

// NewUploadStreamParamsWithContext creates a new UploadStreamParams object
// with the default values initialized, and the ability to set a context for a request
func NewUploadStreamParamsWithContext(ctx context.Context) *UploadStreamParams {
	var ()
	return &UploadStreamParams{

		Context: ctx,
	}
}

// NewUploadStreamParamsWithHTTPClient creates a new UploadStreamParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUploadStreamParamsWithHTTPClient(client *http.Client) *UploadStreamParams {
	var ()
	return &UploadStreamParams{
		HTTPClient: client,
	}
}

/*UploadStreamParams contains all the parameters to send to the API endpoint
for the upload stream operation typically these are written to a http.Request
*/
type UploadStreamParams struct {

	/*XAppendTo
	  Append uploaded data at the end of existing file

	*/
	XAppendTo *string
	/*XPartialTargetBytesize
	  In case of partial upload, the size of the full file as expected at the end of upload.

	*/
	XPartialTargetBytesize *int64
	/*XPartialUpload
	  If the current put is a part of a file. If set, the X-Partial-Target-Bytesize header is required.

	*/
	XPartialUpload *bool
	/*XRenameIfExists
	  Automatically increment filename if it already exists

	*/
	XRenameIfExists *bool
	/*Path
	  Workspace id or alias + full path to the node, e.g. "/my-files/path/to/node"

	*/
	Path string
	/*Raw
	  binary data

	*/
	Raw models.InputStream

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the upload stream params
func (o *UploadStreamParams) WithTimeout(timeout time.Duration) *UploadStreamParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upload stream params
func (o *UploadStreamParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upload stream params
func (o *UploadStreamParams) WithContext(ctx context.Context) *UploadStreamParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upload stream params
func (o *UploadStreamParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upload stream params
func (o *UploadStreamParams) WithHTTPClient(client *http.Client) *UploadStreamParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upload stream params
func (o *UploadStreamParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAppendTo adds the xAppendTo to the upload stream params
func (o *UploadStreamParams) WithXAppendTo(xAppendTo *string) *UploadStreamParams {
	o.SetXAppendTo(xAppendTo)
	return o
}

// SetXAppendTo adds the xAppendTo to the upload stream params
func (o *UploadStreamParams) SetXAppendTo(xAppendTo *string) {
	o.XAppendTo = xAppendTo
}

// WithXPartialTargetBytesize adds the xPartialTargetBytesize to the upload stream params
func (o *UploadStreamParams) WithXPartialTargetBytesize(xPartialTargetBytesize *int64) *UploadStreamParams {
	o.SetXPartialTargetBytesize(xPartialTargetBytesize)
	return o
}

// SetXPartialTargetBytesize adds the xPartialTargetBytesize to the upload stream params
func (o *UploadStreamParams) SetXPartialTargetBytesize(xPartialTargetBytesize *int64) {
	o.XPartialTargetBytesize = xPartialTargetBytesize
}

// WithXPartialUpload adds the xPartialUpload to the upload stream params
func (o *UploadStreamParams) WithXPartialUpload(xPartialUpload *bool) *UploadStreamParams {
	o.SetXPartialUpload(xPartialUpload)
	return o
}

// SetXPartialUpload adds the xPartialUpload to the upload stream params
func (o *UploadStreamParams) SetXPartialUpload(xPartialUpload *bool) {
	o.XPartialUpload = xPartialUpload
}

// WithXRenameIfExists adds the xRenameIfExists to the upload stream params
func (o *UploadStreamParams) WithXRenameIfExists(xRenameIfExists *bool) *UploadStreamParams {
	o.SetXRenameIfExists(xRenameIfExists)
	return o
}

// SetXRenameIfExists adds the xRenameIfExists to the upload stream params
func (o *UploadStreamParams) SetXRenameIfExists(xRenameIfExists *bool) {
	o.XRenameIfExists = xRenameIfExists
}

// WithPath adds the path to the upload stream params
func (o *UploadStreamParams) WithPath(path string) *UploadStreamParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the upload stream params
func (o *UploadStreamParams) SetPath(path string) {
	o.Path = path
}

// WithRaw adds the raw to the upload stream params
func (o *UploadStreamParams) WithRaw(raw models.InputStream) *UploadStreamParams {
	o.SetRaw(raw)
	return o
}

// SetRaw adds the raw to the upload stream params
func (o *UploadStreamParams) SetRaw(raw models.InputStream) {
	o.Raw = raw
}

// WriteToRequest writes these params to a swagger request
func (o *UploadStreamParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XAppendTo != nil {

		// header param X-Append-To
		if err := r.SetHeaderParam("X-Append-To", *o.XAppendTo); err != nil {
			return err
		}

	}

	if o.XPartialTargetBytesize != nil {

		// header param X-Partial-Target-Bytesize
		if err := r.SetHeaderParam("X-Partial-Target-Bytesize", swag.FormatInt64(*o.XPartialTargetBytesize)); err != nil {
			return err
		}

	}

	if o.XPartialUpload != nil {

		// header param X-Partial-Upload
		if err := r.SetHeaderParam("X-Partial-Upload", swag.FormatBool(*o.XPartialUpload)); err != nil {
			return err
		}

	}

	if o.XRenameIfExists != nil {

		// header param X-Rename-If-Exists
		if err := r.SetHeaderParam("X-Rename-If-Exists", swag.FormatBool(*o.XRenameIfExists)); err != nil {
			return err
		}

	}

	// path param path
	if err := r.SetPathParam("path", o.Path); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.Raw); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
