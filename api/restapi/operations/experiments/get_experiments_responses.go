// Code generated by go-swagger; DO NOT EDIT.

package experiments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	lhttp "github.infra.cloudera.com/CAI/AmpRagMonitoring/pkg/http"
)

// GetExperimentsOKCode is the HTTP code returned for type GetExperimentsOK
const GetExperimentsOKCode int = 200

/*
GetExperimentsOK success

swagger:response getExperimentsOK
*/
type GetExperimentsOK struct {

	/*
	  In: Body
	*/
	Payload []string `json:"body,omitempty"`
}

// NewGetExperimentsOK creates GetExperimentsOK with default headers values

func NewGetExperimentsOK() *GetExperimentsOK {

	return &GetExperimentsOK{}
}

// WithPayload adds the payload to the get experiments o k response
func (o *GetExperimentsOK) WithPayload(payload []string) *GetExperimentsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get experiments o k response
func (o *GetExperimentsOK) SetPayload(payload []string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetExperimentsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]string, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetExperimentsBadRequestCode is the HTTP code returned for type GetExperimentsBadRequest
const GetExperimentsBadRequestCode int = 400

/*
GetExperimentsBadRequest bad request

swagger:response getExperimentsBadRequest
*/
type GetExperimentsBadRequest struct {
}

// NewGetExperimentsBadRequest creates GetExperimentsBadRequest with default headers values

func NewGetExperimentsBadRequest() *lhttp.HttpError {
	return &lhttp.HttpError{
		Code: 400,
	}
}

// WriteResponse to the client
func (o *GetExperimentsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetExperimentsInternalServerErrorCode is the HTTP code returned for type GetExperimentsInternalServerError
const GetExperimentsInternalServerErrorCode int = 500

/*
GetExperimentsInternalServerError internal service error

swagger:response getExperimentsInternalServerError
*/
type GetExperimentsInternalServerError struct {
}

// NewGetExperimentsInternalServerError creates GetExperimentsInternalServerError with default headers values

func NewGetExperimentsInternalServerError() *lhttp.HttpError {
	return &lhttp.HttpError{
		Code: 500,
	}
}

// WriteResponse to the client
func (o *GetExperimentsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}