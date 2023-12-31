// Code generated by go-swagger; DO NOT EDIT.

package parser_interface

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TxParserGetTransactionsHandlerFunc turns a function with the right signature into a tx parser get transactions handler
type TxParserGetTransactionsHandlerFunc func(TxParserGetTransactionsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn TxParserGetTransactionsHandlerFunc) Handle(params TxParserGetTransactionsParams) middleware.Responder {
	return fn(params)
}

// TxParserGetTransactionsHandler interface for that can handle valid tx parser get transactions params
type TxParserGetTransactionsHandler interface {
	Handle(TxParserGetTransactionsParams) middleware.Responder
}

// NewTxParserGetTransactions creates a new http.Handler for the tx parser get transactions operation
func NewTxParserGetTransactions(ctx *middleware.Context, handler TxParserGetTransactionsHandler) *TxParserGetTransactions {
	return &TxParserGetTransactions{Context: ctx, Handler: handler}
}

/*
	TxParserGetTransactions swagger:route GET /v1/tx-parser/getTransactions Parser interface txParserGetTransactions

List of inbound or outbound transactions for an address
*/
type TxParserGetTransactions struct {
	Context *middleware.Context
	Handler TxParserGetTransactionsHandler
}

func (o *TxParserGetTransactions) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewTxParserGetTransactionsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// TxParserGetTransactionsDefaultBody tx parser get transactions default body
//
// swagger:model TxParserGetTransactionsDefaultBody
type TxParserGetTransactionsDefaultBody struct {

	// Transactions list
	Transactions []*TxParserGetTransactionsDefaultBodyTransactionsItems0 `json:"transactions"`
}

// Validate validates this tx parser get transactions default body
func (o *TxParserGetTransactionsDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateTransactions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TxParserGetTransactionsDefaultBody) validateTransactions(formats strfmt.Registry) error {
	if swag.IsZero(o.Transactions) { // not required
		return nil
	}

	for i := 0; i < len(o.Transactions); i++ {
		if swag.IsZero(o.Transactions[i]) { // not required
			continue
		}

		if o.Transactions[i] != nil {
			if err := o.Transactions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("txParser_getTransactions default" + "." + "transactions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("txParser_getTransactions default" + "." + "transactions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this tx parser get transactions default body based on the context it is used
func (o *TxParserGetTransactionsDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateTransactions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TxParserGetTransactionsDefaultBody) contextValidateTransactions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Transactions); i++ {

		if o.Transactions[i] != nil {

			if swag.IsZero(o.Transactions[i]) { // not required
				return nil
			}

			if err := o.Transactions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("txParser_getTransactions default" + "." + "transactions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("txParser_getTransactions default" + "." + "transactions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *TxParserGetTransactionsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TxParserGetTransactionsDefaultBody) UnmarshalBinary(b []byte) error {
	var res TxParserGetTransactionsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// TxParserGetTransactionsDefaultBodyTransactionsItems0 tx parser get transactions default body transactions items0
//
// swagger:model TxParserGetTransactionsDefaultBodyTransactionsItems0
type TxParserGetTransactionsDefaultBodyTransactionsItems0 struct {

	// 32 Bytes - hash of the block where this transaction was in. null when its pending.
	BlockHash string `json:"blockHash,omitempty"`

	// block number where this transaction was in. null when its pending.
	BlockNumber string `json:"blockNumber,omitempty"`

	// 20 Bytes - address of the sender.
	From string `json:"from,omitempty"`

	// gas provided by the sender.
	Gas string `json:"gas,omitempty"`

	// gas price provided by the sender in Wei.
	GasPrice string `json:"gasPrice,omitempty"`

	// 32 Bytes - hash of the transaction.
	Hash string `json:"hash,omitempty"`

	// the data send along with the transaction.
	Input string `json:"input,omitempty"`

	// the number of transactions made by the sender prior to this one.
	Nonce string `json:"nonce,omitempty"`

	// ECDSA signature r
	R string `json:"r,omitempty"`

	// ECDSA signature s
	S string `json:"s,omitempty"`

	// 20 Bytes - address of the receiver. null when its a contract creation transaction.
	To string `json:"to,omitempty"`

	// integer of the transactions index position in the block. null when its pending.
	TransactionIndex string `json:"transactionIndex,omitempty"`

	// ECDSA recovery id
	V string `json:"v,omitempty"`

	// value transferred in Wei.
	Value string `json:"value,omitempty"`
}

// Validate validates this tx parser get transactions default body transactions items0
func (o *TxParserGetTransactionsDefaultBodyTransactionsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this tx parser get transactions default body transactions items0 based on context it is used
func (o *TxParserGetTransactionsDefaultBodyTransactionsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *TxParserGetTransactionsDefaultBodyTransactionsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TxParserGetTransactionsDefaultBodyTransactionsItems0) UnmarshalBinary(b []byte) error {
	var res TxParserGetTransactionsDefaultBodyTransactionsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
