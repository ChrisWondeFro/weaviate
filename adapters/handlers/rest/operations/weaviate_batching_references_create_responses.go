//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2019 Weaviate. All rights reserved.
//  LICENSE: https://github.com/semi-technologies/weaviate/blob/develop/LICENSE.md
//  DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// WeaviateBatchingReferencesCreateOKCode is the HTTP code returned for type WeaviateBatchingReferencesCreateOK
const WeaviateBatchingReferencesCreateOKCode int = 200

/*WeaviateBatchingReferencesCreateOK Request Successful. Warning: A successful request does not guarantuee that every batched reference was successfully created. Inspect the response body to see which references succeeded and which failed.

swagger:response weaviateBatchingReferencesCreateOK
*/
type WeaviateBatchingReferencesCreateOK struct {

	/*
	  In: Body
	*/
	Payload []*models.BatchReferenceResponse `json:"body,omitempty"`
}

// NewWeaviateBatchingReferencesCreateOK creates WeaviateBatchingReferencesCreateOK with default headers values
func NewWeaviateBatchingReferencesCreateOK() *WeaviateBatchingReferencesCreateOK {

	return &WeaviateBatchingReferencesCreateOK{}
}

// WithPayload adds the payload to the weaviate batching references create o k response
func (o *WeaviateBatchingReferencesCreateOK) WithPayload(payload []*models.BatchReferenceResponse) *WeaviateBatchingReferencesCreateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate batching references create o k response
func (o *WeaviateBatchingReferencesCreateOK) SetPayload(payload []*models.BatchReferenceResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateBatchingReferencesCreateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.BatchReferenceResponse, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// WeaviateBatchingReferencesCreateUnauthorizedCode is the HTTP code returned for type WeaviateBatchingReferencesCreateUnauthorized
const WeaviateBatchingReferencesCreateUnauthorizedCode int = 401

/*WeaviateBatchingReferencesCreateUnauthorized Unauthorized or invalid credentials.

swagger:response weaviateBatchingReferencesCreateUnauthorized
*/
type WeaviateBatchingReferencesCreateUnauthorized struct {
}

// NewWeaviateBatchingReferencesCreateUnauthorized creates WeaviateBatchingReferencesCreateUnauthorized with default headers values
func NewWeaviateBatchingReferencesCreateUnauthorized() *WeaviateBatchingReferencesCreateUnauthorized {

	return &WeaviateBatchingReferencesCreateUnauthorized{}
}

// WriteResponse to the client
func (o *WeaviateBatchingReferencesCreateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// WeaviateBatchingReferencesCreateForbiddenCode is the HTTP code returned for type WeaviateBatchingReferencesCreateForbidden
const WeaviateBatchingReferencesCreateForbiddenCode int = 403

/*WeaviateBatchingReferencesCreateForbidden Forbidden

swagger:response weaviateBatchingReferencesCreateForbidden
*/
type WeaviateBatchingReferencesCreateForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateBatchingReferencesCreateForbidden creates WeaviateBatchingReferencesCreateForbidden with default headers values
func NewWeaviateBatchingReferencesCreateForbidden() *WeaviateBatchingReferencesCreateForbidden {

	return &WeaviateBatchingReferencesCreateForbidden{}
}

// WithPayload adds the payload to the weaviate batching references create forbidden response
func (o *WeaviateBatchingReferencesCreateForbidden) WithPayload(payload *models.ErrorResponse) *WeaviateBatchingReferencesCreateForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate batching references create forbidden response
func (o *WeaviateBatchingReferencesCreateForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateBatchingReferencesCreateForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateBatchingReferencesCreateUnprocessableEntityCode is the HTTP code returned for type WeaviateBatchingReferencesCreateUnprocessableEntity
const WeaviateBatchingReferencesCreateUnprocessableEntityCode int = 422

/*WeaviateBatchingReferencesCreateUnprocessableEntity Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?

swagger:response weaviateBatchingReferencesCreateUnprocessableEntity
*/
type WeaviateBatchingReferencesCreateUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateBatchingReferencesCreateUnprocessableEntity creates WeaviateBatchingReferencesCreateUnprocessableEntity with default headers values
func NewWeaviateBatchingReferencesCreateUnprocessableEntity() *WeaviateBatchingReferencesCreateUnprocessableEntity {

	return &WeaviateBatchingReferencesCreateUnprocessableEntity{}
}

// WithPayload adds the payload to the weaviate batching references create unprocessable entity response
func (o *WeaviateBatchingReferencesCreateUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *WeaviateBatchingReferencesCreateUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate batching references create unprocessable entity response
func (o *WeaviateBatchingReferencesCreateUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateBatchingReferencesCreateUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateBatchingReferencesCreateInternalServerErrorCode is the HTTP code returned for type WeaviateBatchingReferencesCreateInternalServerError
const WeaviateBatchingReferencesCreateInternalServerErrorCode int = 500

/*WeaviateBatchingReferencesCreateInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response weaviateBatchingReferencesCreateInternalServerError
*/
type WeaviateBatchingReferencesCreateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateBatchingReferencesCreateInternalServerError creates WeaviateBatchingReferencesCreateInternalServerError with default headers values
func NewWeaviateBatchingReferencesCreateInternalServerError() *WeaviateBatchingReferencesCreateInternalServerError {

	return &WeaviateBatchingReferencesCreateInternalServerError{}
}

// WithPayload adds the payload to the weaviate batching references create internal server error response
func (o *WeaviateBatchingReferencesCreateInternalServerError) WithPayload(payload *models.ErrorResponse) *WeaviateBatchingReferencesCreateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate batching references create internal server error response
func (o *WeaviateBatchingReferencesCreateInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateBatchingReferencesCreateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}