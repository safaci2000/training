// This file is safe to edit. Once it exists it will not be overwritten

package server

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/tetrateio/training/samples/modernbank/microservices/account/pkg/server/restapi"
	"github.com/tetrateio/training/samples/modernbank/microservices/account/pkg/server/restapi/accounts"
	"github.com/tetrateio/training/samples/modernbank/microservices/account/pkg/store"
)

var accountStore store.Interface

//go:generate swagger generate server --target ../../../account --name Account --spec ../../../../swagger/account.yaml --api-package restapi --model-package pkg/model --server-package pkg/server

func configureFlags(api *restapi.AccountAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *restapi.AccountAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	accountStore = store.NewInMemory()

	api.AccountsCreateAccountHandler = accounts.CreateAccountHandlerFunc(func(params accounts.CreateAccountParams) middleware.Responder {
		res, err := accountStore.Create(params.Owner)
		if err != nil {
			return accounts.NewCreateAccountInternalServerError()
		}
		return accounts.NewCreateAccountCreated().WithPayload(res)
	})
	api.AccountsDeleteAccountHandler = accounts.DeleteAccountHandlerFunc(func(params accounts.DeleteAccountParams) middleware.Responder {
		if err := accountStore.Delete(params.Owner, params.Number); err != nil {
			if _, ok := err.(*store.NotFound); ok {
				return accounts.NewDeleteAccountNotFound()
			}
			return accounts.NewDeleteAccountInternalServerError()
		}
		return accounts.NewDeleteAccountOK()
	})
	api.AccountsGetAccountByNumberHandler = accounts.GetAccountByNumberHandlerFunc(func(params accounts.GetAccountByNumberParams) middleware.Responder {
		res, err := accountStore.Get(params.Owner, params.Number)
		if err != nil {
			if _, ok := err.(*store.NotFound); ok {
				return accounts.NewGetAccountByNumberNotFound()
			}
			return accounts.NewGetAccountByNumberInternalServerError()
		}
		return accounts.NewGetAccountByNumberOK().WithPayload(res)
	})
	api.AccountsListAccountsHandler = accounts.ListAccountsHandlerFunc(func(params accounts.ListAccountsParams) middleware.Responder {
		res, err := accountStore.List(params.Owner)
		if err != nil {
			if _, ok := err.(*store.NotFound); ok {
				return accounts.NewListAccountsNotFound()
			}
			return accounts.NewListAccountsInternalServerError()
		}
		return accounts.NewListAccountsOK().WithPayload(res)
	})
	api.AccountsChangeBalanceHandler = accounts.ChangeBalanceHandlerFunc(func(params accounts.ChangeBalanceParams) middleware.Responder {
		if err := accountStore.UpdateBalance(params.Number, params.Delta); err != nil {
			if _, ok := err.(*store.NotFound); ok {
				return accounts.NewChangeBalanceNotFound()
			}
			return accounts.NewChangeBalanceInternalServerError()
		}
		return accounts.NewChangeBalanceOK()
	})

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
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
