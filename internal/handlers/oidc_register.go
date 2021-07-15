package handlers

import (
	"github.com/fasthttp/router"

	"github.com/authelia/authelia/internal/middlewares"
)

// RegisterOIDC registers the handlers with the fasthttp *router.Router. TODO: Add paths for UserInfo, Flush, Logout.
func RegisterOIDC(router *router.Router, middleware middlewares.RequestHandlerBridge) {
	// TODO: Add OPTIONS handler.
	router.GET("/.well-known/openid-configuration", middleware(oidcWellKnown))

	router.GET(oidcConsentPath, middleware(oidcConsent))

	router.POST(oidcConsentPath, middleware(oidcConsentPOST))

	router.GET(oidcJWKsPath, middleware(oidcJWKs))

	router.GET(oidcAuthorizationPath, middleware(middlewares.NewHTTPToAutheliaHandlerAdaptor(oidcAuthorization)))

	// TODO: Add OPTIONS handler.
	router.POST(oidcTokenPath, middleware(middlewares.NewHTTPToAutheliaHandlerAdaptor(oidcToken)))

	router.POST(oidcIntrospectionPath, middleware(middlewares.NewHTTPToAutheliaHandlerAdaptor(oidcIntrospection)))

	router.GET(oidcUserinfoPath, middleware(middlewares.NewHTTPToAutheliaHandlerAdaptor(oidcUserinfo)))
	router.POST(oidcUserinfoPath, middleware(middlewares.NewHTTPToAutheliaHandlerAdaptor(oidcUserinfo)))

	// TODO: Add OPTIONS handler.
	router.POST(oidcRevocationPath, middleware(middlewares.NewHTTPToAutheliaHandlerAdaptor(oidcRevocation)))
}
