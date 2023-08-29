// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/pullya/tx_parcer/internal/app/repository"
	"github.com/pullya/tx_parcer/internal/app/service"
	"github.com/pullya/tx_parcer/restapi/operations"
	"github.com/pullya/tx_parcer/restapi/operations/parser_interface"
)

type TxParser struct {
	TxParserService service.TxParserService
}

func NewTxParser(serv service.TxParserService) TxParser {
	return TxParser{
		TxParserService: serv,
	}
}

//go:generate swagger generate server --target ../../tx_parcer --name TxParser --spec ../internal/doc/tx_parser.swagger.json --principal interface{}

func configureFlags(api *operations.TxParserAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.TxParserAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	tpr := repository.NewRepository()
	tps := service.NewTxParserService(tpr)
	tp := NewTxParser(tps)

	api.ParserInterfaceTxParserGetCurrentBlockHandler = parser_interface.TxParserGetCurrentBlockHandlerFunc(func(params parser_interface.TxParserGetCurrentBlockParams) middleware.Responder {
		blockNumber := tp.TxParserService.GetCurrentBlock()
		return parser_interface.NewTxParserGetCurrentBlockOK().WithPayload(&parser_interface.TxParserGetCurrentBlockOKBody{BlockNumber: strconv.Itoa(blockNumber)})
	})

	api.ParserInterfaceTxParserGetTransactionsHandler = parser_interface.TxParserGetTransactionsHandlerFunc(func(params parser_interface.TxParserGetTransactionsParams) middleware.Responder {
		trans := tp.TxParserService.GetTransactions(params.Address)
		response := make([]*parser_interface.TxParserGetTransactionsOKBodyTransactionsItems0, 0, len(trans))
		for _, tran := range trans {
			response = append(response, tran.AsResponse())
		}
		return parser_interface.NewTxParserGetTransactionsOK().WithPayload(&parser_interface.TxParserGetTransactionsOKBody{Transactions: response})
	})

	api.ParserInterfaceTxParserSubscribeHandler = parser_interface.TxParserSubscribeHandlerFunc(func(params parser_interface.TxParserSubscribeParams) middleware.Responder {
		result := tp.TxParserService.Subscribe(params.Address)
		return parser_interface.NewTxParserSubscribeOK().WithPayload(&parser_interface.TxParserSubscribeOKBody{Success: result})

	})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
