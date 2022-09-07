//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package backups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new backups API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for backups API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	BackupsCreate(params *BackupsCreateParams, authInfo runtime.ClientAuthInfoWriter) (*BackupsCreateOK, error)

	BackupsCreateStatus(params *BackupsCreateStatusParams, authInfo runtime.ClientAuthInfoWriter) (*BackupsCreateStatusOK, error)

	BackupsRestore(params *BackupsRestoreParams, authInfo runtime.ClientAuthInfoWriter) (*BackupsRestoreOK, error)

	BackupsRestoreStatus(params *BackupsRestoreStatusParams, authInfo runtime.ClientAuthInfoWriter) (*BackupsRestoreStatusOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  BackupsCreate Starts a process of creating a backup for a set of classes
*/
func (a *Client) BackupsCreate(params *BackupsCreateParams, authInfo runtime.ClientAuthInfoWriter) (*BackupsCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupsCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "backups.create",
		Method:             "POST",
		PathPattern:        "/backups/{backend}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BackupsCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BackupsCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for backups.create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  BackupsCreateStatus Returns status of backup creation attempt for a set of classes
*/
func (a *Client) BackupsCreateStatus(params *BackupsCreateStatusParams, authInfo runtime.ClientAuthInfoWriter) (*BackupsCreateStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupsCreateStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "backups.create.status",
		Method:             "GET",
		PathPattern:        "/backups/{backend}/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BackupsCreateStatusReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BackupsCreateStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for backups.create.status: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  BackupsRestore Starts a process of restoring a backup for a set of classes
*/
func (a *Client) BackupsRestore(params *BackupsRestoreParams, authInfo runtime.ClientAuthInfoWriter) (*BackupsRestoreOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupsRestoreParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "backups.restore",
		Method:             "POST",
		PathPattern:        "/backups/{backend}/{id}/restore",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BackupsRestoreReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BackupsRestoreOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for backups.restore: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  BackupsRestoreStatus Returns status of a backup restoration attempt for a set of classes
*/
func (a *Client) BackupsRestoreStatus(params *BackupsRestoreStatusParams, authInfo runtime.ClientAuthInfoWriter) (*BackupsRestoreStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupsRestoreStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "backups.restore.status",
		Method:             "GET",
		PathPattern:        "/backups/{backend}/{id}/restore",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BackupsRestoreStatusReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BackupsRestoreStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for backups.restore.status: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}