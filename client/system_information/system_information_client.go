// Code generated by go-swagger; DO NOT EDIT.

package system_information

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new system information API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for system information API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetSystemInformation gets s e m p API version and platform information

Gets SEMP API version and platform information.

A SEMP client authorized with a minimum access scope/level of "global/none" is required to perform this operation.

This has been available since 2.1.0.

This has been deprecated since 2.2.0.
*/
func (a *Client) GetSystemInformation(params *GetSystemInformationParams, authInfo runtime.ClientAuthInfoWriter) (*GetSystemInformationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSystemInformationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSystemInformation",
		Method:             "GET",
		PathPattern:        "/systemInformation",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSystemInformationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSystemInformationOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
