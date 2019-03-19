// Code generated by go-swagger; DO NOT EDIT.

package transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	model "github.com/tetrateio/training/samples/modernbank/microservices/transaction-log/pkg/model"
)

// GetTransactionSentOKCode is the HTTP code returned for type GetTransactionSentOK
const GetTransactionSentOKCode int = 200

/*GetTransactionSentOK Success!

swagger:response getTransactionSentOK
*/
type GetTransactionSentOK struct {
	/*Version of the microservice that responded

	 */
	Version string `json:"version"`

	/*
	  In: Body
	*/
	Payload *model.Transaction `json:"body,omitempty"`
}

// NewGetTransactionSentOK creates GetTransactionSentOK with default headers values
func NewGetTransactionSentOK() *GetTransactionSentOK {

	return &GetTransactionSentOK{}
}

// WithVersion adds the version to the get transaction sent o k response
func (o *GetTransactionSentOK) WithVersion(version string) *GetTransactionSentOK {
	o.Version = version
	return o
}

// SetVersion sets the version to the get transaction sent o k response
func (o *GetTransactionSentOK) SetVersion(version string) {
	o.Version = version
}

// WithPayload adds the payload to the get transaction sent o k response
func (o *GetTransactionSentOK) WithPayload(payload *model.Transaction) *GetTransactionSentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get transaction sent o k response
func (o *GetTransactionSentOK) SetPayload(payload *model.Transaction) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTransactionSentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header version

	version := o.Version
	if version != "" {
		rw.Header().Set("version", version)
	}

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTransactionSentNotFoundCode is the HTTP code returned for type GetTransactionSentNotFound
const GetTransactionSentNotFoundCode int = 404

/*GetTransactionSentNotFound Transaction not found

swagger:response getTransactionSentNotFound
*/
type GetTransactionSentNotFound struct {
	/*Version of the microservice that responded

	 */
	Version string `json:"version"`
}

// NewGetTransactionSentNotFound creates GetTransactionSentNotFound with default headers values
func NewGetTransactionSentNotFound() *GetTransactionSentNotFound {

	return &GetTransactionSentNotFound{}
}

// WithVersion adds the version to the get transaction sent not found response
func (o *GetTransactionSentNotFound) WithVersion(version string) *GetTransactionSentNotFound {
	o.Version = version
	return o
}

// SetVersion sets the version to the get transaction sent not found response
func (o *GetTransactionSentNotFound) SetVersion(version string) {
	o.Version = version
}

// WriteResponse to the client
func (o *GetTransactionSentNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header version

	version := o.Version
	if version != "" {
		rw.Header().Set("version", version)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetTransactionSentInternalServerErrorCode is the HTTP code returned for type GetTransactionSentInternalServerError
const GetTransactionSentInternalServerErrorCode int = 500

/*GetTransactionSentInternalServerError Internal server error

swagger:response getTransactionSentInternalServerError
*/
type GetTransactionSentInternalServerError struct {
	/*Version of the microservice that responded

	 */
	Version string `json:"version"`
}

// NewGetTransactionSentInternalServerError creates GetTransactionSentInternalServerError with default headers values
func NewGetTransactionSentInternalServerError() *GetTransactionSentInternalServerError {

	return &GetTransactionSentInternalServerError{}
}

// WithVersion adds the version to the get transaction sent internal server error response
func (o *GetTransactionSentInternalServerError) WithVersion(version string) *GetTransactionSentInternalServerError {
	o.Version = version
	return o
}

// SetVersion sets the version to the get transaction sent internal server error response
func (o *GetTransactionSentInternalServerError) SetVersion(version string) {
	o.Version = version
}

// WriteResponse to the client
func (o *GetTransactionSentInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header version

	version := o.Version
	if version != "" {
		rw.Header().Set("version", version)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
