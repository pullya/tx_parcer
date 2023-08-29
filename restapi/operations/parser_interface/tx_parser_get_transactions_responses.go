// Code generated by go-swagger; DO NOT EDIT.

package parser_interface

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pullya/tx_parcer/models"
)

// TxParserGetTransactionsOKCode is the HTTP code returned for type TxParserGetTransactionsOK
const TxParserGetTransactionsOKCode int = 200

/*
TxParserGetTransactionsOK Successful response

swagger:response txParserGetTransactionsOK
*/
type TxParserGetTransactionsOK struct {

	/*
	  In: Body
	*/
	Payload *TxParserGetTransactionsOKBody `json:"body,omitempty"`
}

// NewTxParserGetTransactionsOK creates TxParserGetTransactionsOK with default headers values
func NewTxParserGetTransactionsOK() *TxParserGetTransactionsOK {

	return &TxParserGetTransactionsOK{}
}

// WithPayload adds the payload to the tx parser get transactions o k response
func (o *TxParserGetTransactionsOK) WithPayload(payload *TxParserGetTransactionsOKBody) *TxParserGetTransactionsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the tx parser get transactions o k response
func (o *TxParserGetTransactionsOK) SetPayload(payload *TxParserGetTransactionsOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TxParserGetTransactionsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
TxParserGetTransactionsDefault Unsuccessful response

swagger:response txParserGetTransactionsDefault
*/
type TxParserGetTransactionsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewTxParserGetTransactionsDefault creates TxParserGetTransactionsDefault with default headers values
func NewTxParserGetTransactionsDefault(code int) *TxParserGetTransactionsDefault {
	if code <= 0 {
		code = 500
	}

	return &TxParserGetTransactionsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the tx parser get transactions default response
func (o *TxParserGetTransactionsDefault) WithStatusCode(code int) *TxParserGetTransactionsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the tx parser get transactions default response
func (o *TxParserGetTransactionsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the tx parser get transactions default response
func (o *TxParserGetTransactionsDefault) WithPayload(payload *models.ErrorResponse) *TxParserGetTransactionsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the tx parser get transactions default response
func (o *TxParserGetTransactionsDefault) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TxParserGetTransactionsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
