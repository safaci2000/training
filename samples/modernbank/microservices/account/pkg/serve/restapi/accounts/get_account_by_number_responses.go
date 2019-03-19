// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	model "github.com/tetrateio/training/samples/modernbank/microservices/account/pkg/model"
)

// GetAccountByNumberOKCode is the HTTP code returned for type GetAccountByNumberOK
const GetAccountByNumberOKCode int = 200

/*GetAccountByNumberOK Success!

swagger:response getAccountByNumberOK
*/
type GetAccountByNumberOK struct {
	/*Version of the microservice that responded

	 */
	Version string `json:"version"`

	/*
	  In: Body
	*/
	Payload *model.Account `json:"body,omitempty"`
}

// NewGetAccountByNumberOK creates GetAccountByNumberOK with default headers values
func NewGetAccountByNumberOK() *GetAccountByNumberOK {

	return &GetAccountByNumberOK{}
}

// WithVersion adds the version to the get account by number o k response
func (o *GetAccountByNumberOK) WithVersion(version string) *GetAccountByNumberOK {
	o.Version = version
	return o
}

// SetVersion sets the version to the get account by number o k response
func (o *GetAccountByNumberOK) SetVersion(version string) {
	o.Version = version
}

// WithPayload adds the payload to the get account by number o k response
func (o *GetAccountByNumberOK) WithPayload(payload *model.Account) *GetAccountByNumberOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get account by number o k response
func (o *GetAccountByNumberOK) SetPayload(payload *model.Account) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAccountByNumberOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetAccountByNumberNotFoundCode is the HTTP code returned for type GetAccountByNumberNotFound
const GetAccountByNumberNotFoundCode int = 404

/*GetAccountByNumberNotFound Account not found

swagger:response getAccountByNumberNotFound
*/
type GetAccountByNumberNotFound struct {
	/*Version of the microservice that responded

	 */
	Version string `json:"version"`
}

// NewGetAccountByNumberNotFound creates GetAccountByNumberNotFound with default headers values
func NewGetAccountByNumberNotFound() *GetAccountByNumberNotFound {

	return &GetAccountByNumberNotFound{}
}

// WithVersion adds the version to the get account by number not found response
func (o *GetAccountByNumberNotFound) WithVersion(version string) *GetAccountByNumberNotFound {
	o.Version = version
	return o
}

// SetVersion sets the version to the get account by number not found response
func (o *GetAccountByNumberNotFound) SetVersion(version string) {
	o.Version = version
}

// WriteResponse to the client
func (o *GetAccountByNumberNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header version

	version := o.Version
	if version != "" {
		rw.Header().Set("version", version)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetAccountByNumberInternalServerErrorCode is the HTTP code returned for type GetAccountByNumberInternalServerError
const GetAccountByNumberInternalServerErrorCode int = 500

/*GetAccountByNumberInternalServerError Internal server error

swagger:response getAccountByNumberInternalServerError
*/
type GetAccountByNumberInternalServerError struct {
	/*Version of the microservice that responded

	 */
	Version string `json:"version"`
}

// NewGetAccountByNumberInternalServerError creates GetAccountByNumberInternalServerError with default headers values
func NewGetAccountByNumberInternalServerError() *GetAccountByNumberInternalServerError {

	return &GetAccountByNumberInternalServerError{}
}

// WithVersion adds the version to the get account by number internal server error response
func (o *GetAccountByNumberInternalServerError) WithVersion(version string) *GetAccountByNumberInternalServerError {
	o.Version = version
	return o
}

// SetVersion sets the version to the get account by number internal server error response
func (o *GetAccountByNumberInternalServerError) SetVersion(version string) {
	o.Version = version
}

// WriteResponse to the client
func (o *GetAccountByNumberInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header version

	version := o.Version
	if version != "" {
		rw.Header().Set("version", version)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
