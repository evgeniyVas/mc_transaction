// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mc_transaction/internal/http_client/paysystem/models"
)

// PostAPITransactionsReverseReader is a Reader for the PostAPITransactionsReverse structure.
type PostAPITransactionsReverseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPITransactionsReverseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAPITransactionsReverseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAPITransactionsReverseBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAPITransactionsReverseInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/transactions/reverse] PostAPITransactionsReverse", response, response.Code())
	}
}

// NewPostAPITransactionsReverseOK creates a PostAPITransactionsReverseOK with default headers values
func NewPostAPITransactionsReverseOK() *PostAPITransactionsReverseOK {
	return &PostAPITransactionsReverseOK{}
}

/*
PostAPITransactionsReverseOK describes a response with status code 200, with default header values.

ОК
*/
type PostAPITransactionsReverseOK struct {
}

// IsSuccess returns true when this post Api transactions reverse o k response has a 2xx status code
func (o *PostAPITransactionsReverseOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post Api transactions reverse o k response has a 3xx status code
func (o *PostAPITransactionsReverseOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Api transactions reverse o k response has a 4xx status code
func (o *PostAPITransactionsReverseOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post Api transactions reverse o k response has a 5xx status code
func (o *PostAPITransactionsReverseOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post Api transactions reverse o k response a status code equal to that given
func (o *PostAPITransactionsReverseOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post Api transactions reverse o k response
func (o *PostAPITransactionsReverseOK) Code() int {
	return 200
}

func (o *PostAPITransactionsReverseOK) Error() string {
	return fmt.Sprintf("[POST /api/transactions/reverse][%d] postApiTransactionsReverseOK ", 200)
}

func (o *PostAPITransactionsReverseOK) String() string {
	return fmt.Sprintf("[POST /api/transactions/reverse][%d] postApiTransactionsReverseOK ", 200)
}

func (o *PostAPITransactionsReverseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAPITransactionsReverseBadRequest creates a PostAPITransactionsReverseBadRequest with default headers values
func NewPostAPITransactionsReverseBadRequest() *PostAPITransactionsReverseBadRequest {
	return &PostAPITransactionsReverseBadRequest{}
}

/*
PostAPITransactionsReverseBadRequest describes a response with status code 400, with default header values.

Ошибка клиента
*/
type PostAPITransactionsReverseBadRequest struct {
}

// IsSuccess returns true when this post Api transactions reverse bad request response has a 2xx status code
func (o *PostAPITransactionsReverseBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post Api transactions reverse bad request response has a 3xx status code
func (o *PostAPITransactionsReverseBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Api transactions reverse bad request response has a 4xx status code
func (o *PostAPITransactionsReverseBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post Api transactions reverse bad request response has a 5xx status code
func (o *PostAPITransactionsReverseBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post Api transactions reverse bad request response a status code equal to that given
func (o *PostAPITransactionsReverseBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post Api transactions reverse bad request response
func (o *PostAPITransactionsReverseBadRequest) Code() int {
	return 400
}

func (o *PostAPITransactionsReverseBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/transactions/reverse][%d] postApiTransactionsReverseBadRequest ", 400)
}

func (o *PostAPITransactionsReverseBadRequest) String() string {
	return fmt.Sprintf("[POST /api/transactions/reverse][%d] postApiTransactionsReverseBadRequest ", 400)
}

func (o *PostAPITransactionsReverseBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAPITransactionsReverseInternalServerError creates a PostAPITransactionsReverseInternalServerError with default headers values
func NewPostAPITransactionsReverseInternalServerError() *PostAPITransactionsReverseInternalServerError {
	return &PostAPITransactionsReverseInternalServerError{}
}

/*
PostAPITransactionsReverseInternalServerError describes a response with status code 500, with default header values.

Internal error
*/
type PostAPITransactionsReverseInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this post Api transactions reverse internal server error response has a 2xx status code
func (o *PostAPITransactionsReverseInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post Api transactions reverse internal server error response has a 3xx status code
func (o *PostAPITransactionsReverseInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Api transactions reverse internal server error response has a 4xx status code
func (o *PostAPITransactionsReverseInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post Api transactions reverse internal server error response has a 5xx status code
func (o *PostAPITransactionsReverseInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post Api transactions reverse internal server error response a status code equal to that given
func (o *PostAPITransactionsReverseInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post Api transactions reverse internal server error response
func (o *PostAPITransactionsReverseInternalServerError) Code() int {
	return 500
}

func (o *PostAPITransactionsReverseInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/transactions/reverse][%d] postApiTransactionsReverseInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAPITransactionsReverseInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/transactions/reverse][%d] postApiTransactionsReverseInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAPITransactionsReverseInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostAPITransactionsReverseInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}