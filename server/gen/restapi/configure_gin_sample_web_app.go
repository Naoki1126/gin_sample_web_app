// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"gin_sample_web_app/handler"
	"gin_sample_web_app/server"
	"gin_sample_web_app/server/gen/restapi/gin_sample_web_app"
)

//go:generate swagger generate server --target ../../gen --name GinSampleWebApp --spec ../../../swagger/swagger.yaml --api-package gin_sample_web_app --principal interface{}

func configureFlags(api *gin_sample_web_app.GinSampleWebAppAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *gin_sample_web_app.GinSampleWebAppAPI) http.Handler {
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

	api.TxtProducer = runtime.TextProducer()

	api.GetGreetingHandler = gin_sample_web_app.GetGreetingHandlerFunc(server.GetGreeting)

	api.GetUserByIDHandler = gin_sample_web_app.GetUserByIDHandlerFunc(func(params gin_sample_web_app.GetUserByIDParams) middleware.Responder {
		h := handler.GetUserByIDHandler{}
		return h.Handle(params)
	})
	api.PostUserHandler = gin_sample_web_app.PostUserHandlerFunc(func(params gin_sample_web_app.PostUserParams) middleware.Responder {
		h := handler.PostUserHandler{}
		return h.Handle(params)
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
