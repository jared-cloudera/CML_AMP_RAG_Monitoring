// Code generated by go-swagger; DO NOT EDIT.

package runs

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
)

// NewDeleteRunsParams creates a new DeleteRunsParams object
// with the default values initialized.
func NewDeleteRunsParams() *DeleteRunsParams {
	var ()
	return &DeleteRunsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRunsParamsWithTimeout creates a new DeleteRunsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteRunsParamsWithTimeout(timeout time.Duration) *DeleteRunsParams {
	var ()
	return &DeleteRunsParams{

		timeout: timeout,
	}
}

// NewDeleteRunsParamsWithContext creates a new DeleteRunsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteRunsParamsWithContext(ctx context.Context) *DeleteRunsParams {
	var ()
	return &DeleteRunsParams{

		Context: ctx,
	}
}

// NewDeleteRunsParamsWithHTTPClient creates a new DeleteRunsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteRunsParamsWithHTTPClient(client *http.Client) *DeleteRunsParams {
	var ()
	return &DeleteRunsParams{
		HTTPClient: client,
	}
}

/*
DeleteRunsParams contains all the parameters to send to the API endpoint
for the delete runs operation typically these are written to a http.Request
*/
type DeleteRunsParams struct {

	/*ExperimentID
	  The ID of the experiment

	*/
	ExperimentID *string
	/*RunID
	  The ID of the run

	*/
	RunID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

var _ lswagger.SwaggerParams = &DeleteRunsParams{}

func (o *DeleteRunsParams) GetSerializedParams() ([]byte, error) {
	var params = struct {
		ExperimentID *string

		RunID *string
	}{

		ExperimentID: o.ExperimentID,

		RunID: o.RunID,
	}

	return json.Marshal(&params)
}

// WithTimeout adds the timeout to the delete runs params
func (o *DeleteRunsParams) WithTimeout(timeout time.Duration) *DeleteRunsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete runs params
func (o *DeleteRunsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete runs params
func (o *DeleteRunsParams) WithContext(ctx context.Context) *DeleteRunsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete runs params
func (o *DeleteRunsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete runs params
func (o *DeleteRunsParams) WithHTTPClient(client *http.Client) *DeleteRunsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete runs params
func (o *DeleteRunsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExperimentID adds the experimentID to the delete runs params
func (o *DeleteRunsParams) WithExperimentID(experimentID *string) *DeleteRunsParams {
	o.SetExperimentID(experimentID)
	return o
}

// SetExperimentID adds the experimentId to the delete runs params
func (o *DeleteRunsParams) SetExperimentID(experimentID *string) {
	o.ExperimentID = experimentID
}

// WithRunID adds the runID to the delete runs params
func (o *DeleteRunsParams) WithRunID(runID *string) *DeleteRunsParams {
	o.SetRunID(runID)
	return o
}

// SetRunID adds the runId to the delete runs params
func (o *DeleteRunsParams) SetRunID(runID *string) {
	o.RunID = runID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRunsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ExperimentID != nil {

		// query param experiment_id
		var qrExperimentID string
		if o.ExperimentID != nil {
			qrExperimentID = *o.ExperimentID
		}
		qExperimentID := qrExperimentID
		if qExperimentID != "" {
			if err := r.SetQueryParam("experiment_id", qExperimentID); err != nil {
				return err
			}
		}

	}

	if o.RunID != nil {

		// query param run_id
		var qrRunID string
		if o.RunID != nil {
			qrRunID = *o.RunID
		}
		qRunID := qrRunID
		if qRunID != "" {
			if err := r.SetQueryParam("run_id", qRunID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}