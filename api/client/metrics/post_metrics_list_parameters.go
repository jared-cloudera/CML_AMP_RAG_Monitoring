// Code generated by go-swagger; DO NOT EDIT.

package metrics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	lswagger "github.infra.cloudera.com/CAI/AmpRagMonitoring/pkg/swagger"

	strfmt "github.com/go-openapi/strfmt"

	"github.infra.cloudera.com/CAI/AmpRagMonitoring/models"
)

// NewPostMetricsListParams creates a new PostMetricsListParams object
// with the default values initialized.
func NewPostMetricsListParams() *PostMetricsListParams {
	var ()
	return &PostMetricsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostMetricsListParamsWithTimeout creates a new PostMetricsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostMetricsListParamsWithTimeout(timeout time.Duration) *PostMetricsListParams {
	var ()
	return &PostMetricsListParams{

		timeout: timeout,
	}
}

// NewPostMetricsListParamsWithContext creates a new PostMetricsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostMetricsListParamsWithContext(ctx context.Context) *PostMetricsListParams {
	var ()
	return &PostMetricsListParams{

		Context: ctx,
	}
}

// NewPostMetricsListParamsWithHTTPClient creates a new PostMetricsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostMetricsListParamsWithHTTPClient(client *http.Client) *PostMetricsListParams {
	var ()
	return &PostMetricsListParams{
		HTTPClient: client,
	}
}

/*
PostMetricsListParams contains all the parameters to send to the API endpoint
for the post metrics list operation typically these are written to a http.Request
*/
type PostMetricsListParams struct {

	/*Body*/
	Body *models.MetricListFilter

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

var _ lswagger.SwaggerParams = &PostMetricsListParams{}

func (o *PostMetricsListParams) GetSerializedParams() ([]byte, error) {
	var params = struct {
		Body *models.MetricListFilter
	}{

		Body: o.Body,
	}

	return json.Marshal(&params)
}

// WithTimeout adds the timeout to the post metrics list params
func (o *PostMetricsListParams) WithTimeout(timeout time.Duration) *PostMetricsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post metrics list params
func (o *PostMetricsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post metrics list params
func (o *PostMetricsListParams) WithContext(ctx context.Context) *PostMetricsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post metrics list params
func (o *PostMetricsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post metrics list params
func (o *PostMetricsListParams) WithHTTPClient(client *http.Client) *PostMetricsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post metrics list params
func (o *PostMetricsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post metrics list params
func (o *PostMetricsListParams) WithBody(body *models.MetricListFilter) *PostMetricsListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post metrics list params
func (o *PostMetricsListParams) SetBody(body *models.MetricListFilter) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostMetricsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}