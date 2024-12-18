// Code generated by go-swagger; DO NOT EDIT.

package runs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	lhttp "github.infra.cloudera.com/CAI/AmpRagMonitoring/pkg/http"

	"github.infra.cloudera.com/CAI/AmpRagMonitoring/models"
)

// PostRunsListOKCode is the HTTP code returned for type PostRunsListOK
const PostRunsListOKCode int = 200

/*
PostRunsListOK success

swagger:response postRunsListOK
*/
type PostRunsListOK struct {

	/*
	  In: Body
	*/
	Payload []*models.ExperimentRun `json:"body,omitempty"`
}

// NewPostRunsListOK creates PostRunsListOK with default headers values

func NewPostRunsListOK() *PostRunsListOK {

	return &PostRunsListOK{}
}

// WithPayload adds the payload to the post runs list o k response
func (o *PostRunsListOK) WithPayload(payload []*models.ExperimentRun) *PostRunsListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post runs list o k response
func (o *PostRunsListOK) SetPayload(payload []*models.ExperimentRun) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostRunsListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.ExperimentRun, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PostRunsListBadRequestCode is the HTTP code returned for type PostRunsListBadRequest
const PostRunsListBadRequestCode int = 400

/*
PostRunsListBadRequest bad request

swagger:response postRunsListBadRequest
*/
type PostRunsListBadRequest struct {
}

// NewPostRunsListBadRequest creates PostRunsListBadRequest with default headers values

func NewPostRunsListBadRequest() *lhttp.HttpError {
	return &lhttp.HttpError{
		Code: 400,
	}
}

// WriteResponse to the client
func (o *PostRunsListBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// PostRunsListInternalServerErrorCode is the HTTP code returned for type PostRunsListInternalServerError
const PostRunsListInternalServerErrorCode int = 500

/*
PostRunsListInternalServerError internal service error

swagger:response postRunsListInternalServerError
*/
type PostRunsListInternalServerError struct {
}

// NewPostRunsListInternalServerError creates PostRunsListInternalServerError with default headers values

func NewPostRunsListInternalServerError() *lhttp.HttpError {
	return &lhttp.HttpError{
		Code: 500,
	}
}

// WriteResponse to the client
func (o *PostRunsListInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
