package middlewares

import "github.com/valyala/fasthttp"

const (
	headerOrigin                      = "Origin"
	headerAccessControlRequestHeaders = "Access-Control-Request-Headers"
	headerAccessControlRequestMethod  = "Access-Control-Request-Method"

	headerAccessControlAllowOrigin      = "Access-Control-Allow-Origin"
	headerAccessControlAllowHeaders     = "Access-Control-Allow-Headers"
	headerAccessControlAllowMethods     = "Access-Control-Allow-Methods"
	headerAccessControlAllowCredentials = "Access-Control-Allow-Credentials"
	headerVary                          = "Vary"
)

// AutomaticCORSMiddleware automatically adds all relevant CORS headers to a request.
func AutomaticCORSMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		corsOrigin := ctx.Request.Header.Peek(headerOrigin)

		// TODO: Check Origin is protected by Authelia.
		if corsOrigin != nil {
			ctx.Response.Header.SetBytesV(headerAccessControlAllowOrigin, corsOrigin)
			ctx.Response.Header.Set(headerVary, "Accept-Encoding, Origin")
			ctx.Response.Header.Set(headerAccessControlAllowCredentials, "false")
			ctx.Response.Header.SetBytesV(headerAccessControlAllowHeaders, ctx.Request.Header.Peek(headerAccessControlRequestHeaders))
			ctx.Response.Header.SetBytesV(headerAccessControlAllowMethods, ctx.Request.Header.Peek(headerAccessControlRequestMethod))
		}

		next(ctx)
	}
}
