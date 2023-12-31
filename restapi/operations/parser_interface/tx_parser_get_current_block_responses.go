// Code generated by go-swagger; DO NOT EDIT.

package parser_interface

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

/*
TxParserGetCurrentBlockDefault Method response

swagger:response txParserGetCurrentBlockDefault
*/
type TxParserGetCurrentBlockDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *TxParserGetCurrentBlockDefaultBody `json:"body,omitempty"`
}

// NewTxParserGetCurrentBlockDefault creates TxParserGetCurrentBlockDefault with default headers values
func NewTxParserGetCurrentBlockDefault(code int) *TxParserGetCurrentBlockDefault {
	if code <= 0 {
		code = 500
	}

	return &TxParserGetCurrentBlockDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the tx parser get current block default response
func (o *TxParserGetCurrentBlockDefault) WithStatusCode(code int) *TxParserGetCurrentBlockDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the tx parser get current block default response
func (o *TxParserGetCurrentBlockDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the tx parser get current block default response
func (o *TxParserGetCurrentBlockDefault) WithPayload(payload *TxParserGetCurrentBlockDefaultBody) *TxParserGetCurrentBlockDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the tx parser get current block default response
func (o *TxParserGetCurrentBlockDefault) SetPayload(payload *TxParserGetCurrentBlockDefaultBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TxParserGetCurrentBlockDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
