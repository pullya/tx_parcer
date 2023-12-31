// Code generated by go-swagger; DO NOT EDIT.

package parser_interface

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TxParserSubscribeHandlerFunc turns a function with the right signature into a tx parser subscribe handler
type TxParserSubscribeHandlerFunc func(TxParserSubscribeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn TxParserSubscribeHandlerFunc) Handle(params TxParserSubscribeParams) middleware.Responder {
	return fn(params)
}

// TxParserSubscribeHandler interface for that can handle valid tx parser subscribe params
type TxParserSubscribeHandler interface {
	Handle(TxParserSubscribeParams) middleware.Responder
}

// NewTxParserSubscribe creates a new http.Handler for the tx parser subscribe operation
func NewTxParserSubscribe(ctx *middleware.Context, handler TxParserSubscribeHandler) *TxParserSubscribe {
	return &TxParserSubscribe{Context: ctx, Handler: handler}
}

/*
	TxParserSubscribe swagger:route POST /v1/tx-parser/subscribe Parser interface txParserSubscribe

Add address to observer
*/
type TxParserSubscribe struct {
	Context *middleware.Context
	Handler TxParserSubscribeHandler
}

func (o *TxParserSubscribe) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewTxParserSubscribeParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// TxParserSubscribeDefaultBody tx parser subscribe default body
//
// swagger:model TxParserSubscribeDefaultBody
type TxParserSubscribeDefaultBody struct {

	// success
	Success bool `json:"success,omitempty"`
}

// Validate validates this tx parser subscribe default body
func (o *TxParserSubscribeDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this tx parser subscribe default body based on context it is used
func (o *TxParserSubscribeDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *TxParserSubscribeDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TxParserSubscribeDefaultBody) UnmarshalBinary(b []byte) error {
	var res TxParserSubscribeDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
