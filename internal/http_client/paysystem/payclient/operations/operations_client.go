// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetAPITransactionsIDStatus(params *GetAPITransactionsIDStatusParams, opts ...ClientOption) (*GetAPITransactionsIDStatusOK, error)

	PostAPITransactions(params *PostAPITransactionsParams, opts ...ClientOption) (*PostAPITransactionsOK, error)

	PostAPITransactionsReverse(params *PostAPITransactionsReverseParams, opts ...ClientOption) (*PostAPITransactionsReverseOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetAPITransactionsIDStatus get API transactions ID status API
*/
func (a *Client) GetAPITransactionsIDStatus(params *GetAPITransactionsIDStatusParams, opts ...ClientOption) (*GetAPITransactionsIDStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPITransactionsIDStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPITransactionsIDStatus",
		Method:             "GET",
		PathPattern:        "/api/transactions/{id}/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAPITransactionsIDStatusReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAPITransactionsIDStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPITransactionsIDStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostAPITransactions post API transactions API
*/
func (a *Client) PostAPITransactions(params *PostAPITransactionsParams, opts ...ClientOption) (*PostAPITransactionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAPITransactionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostAPITransactions",
		Method:             "POST",
		PathPattern:        "/api/transactions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostAPITransactionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostAPITransactionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostAPITransactions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostAPITransactionsReverse post API transactions reverse API
*/
func (a *Client) PostAPITransactionsReverse(params *PostAPITransactionsReverseParams, opts ...ClientOption) (*PostAPITransactionsReverseOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAPITransactionsReverseParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostAPITransactionsReverse",
		Method:             "POST",
		PathPattern:        "/api/transactions/reverse",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostAPITransactionsReverseReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostAPITransactionsReverseOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostAPITransactionsReverse: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
