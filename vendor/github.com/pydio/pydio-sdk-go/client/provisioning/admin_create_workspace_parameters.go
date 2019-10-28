// Code generated by go-swagger; DO NOT EDIT.

package provisioning

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/pydio/pydio-sdk-go/models"
)

// NewAdminCreateWorkspaceParams creates a new AdminCreateWorkspaceParams object
// with the default values initialized.
func NewAdminCreateWorkspaceParams() *AdminCreateWorkspaceParams {
	var ()
	return &AdminCreateWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminCreateWorkspaceParamsWithTimeout creates a new AdminCreateWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminCreateWorkspaceParamsWithTimeout(timeout time.Duration) *AdminCreateWorkspaceParams {
	var ()
	return &AdminCreateWorkspaceParams{

		timeout: timeout,
	}
}

// NewAdminCreateWorkspaceParamsWithContext creates a new AdminCreateWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewAdminCreateWorkspaceParamsWithContext(ctx context.Context) *AdminCreateWorkspaceParams {
	var ()
	return &AdminCreateWorkspaceParams{

		Context: ctx,
	}
}

// NewAdminCreateWorkspaceParamsWithHTTPClient creates a new AdminCreateWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminCreateWorkspaceParamsWithHTTPClient(client *http.Client) *AdminCreateWorkspaceParams {
	var ()
	return &AdminCreateWorkspaceParams{
		HTTPClient: client,
	}
}

/*AdminCreateWorkspaceParams contains all the parameters to send to the API endpoint
for the admin create workspace operation typically these are written to a http.Request
*/
type AdminCreateWorkspaceParams struct {

	/*Payload
	  Repository details

	*/
	Payload *models.AdminWorkspace

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the admin create workspace params
func (o *AdminCreateWorkspaceParams) WithTimeout(timeout time.Duration) *AdminCreateWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin create workspace params
func (o *AdminCreateWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin create workspace params
func (o *AdminCreateWorkspaceParams) WithContext(ctx context.Context) *AdminCreateWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin create workspace params
func (o *AdminCreateWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin create workspace params
func (o *AdminCreateWorkspaceParams) WithHTTPClient(client *http.Client) *AdminCreateWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin create workspace params
func (o *AdminCreateWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPayload adds the payload to the admin create workspace params
func (o *AdminCreateWorkspaceParams) WithPayload(payload *models.AdminWorkspace) *AdminCreateWorkspaceParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the admin create workspace params
func (o *AdminCreateWorkspaceParams) SetPayload(payload *models.AdminWorkspace) {
	o.Payload = payload
}

// WriteToRequest writes these params to a swagger request
func (o *AdminCreateWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Payload != nil {
		if err := r.SetBodyParam(o.Payload); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
