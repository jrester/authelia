package middlewares

import (
	"fmt"
	"net/url"

	"github.com/valyala/fasthttp"

	"github.com/authelia/authelia/internal/utils"
)

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
func AutomaticCORSMiddleware(next RequestHandler) RequestHandler {
	return func(ctx *AutheliaCtx) {
		if ctx.IsXHR() || !ctx.AcceptsMIME("text/html") {
			host, err := ctx.ForwardedProtoHost()
			if err != nil {
				ctx.ReplyBadRequest()
				return
			}

			ctx.SpecialRedirect(fmt.Sprintf("%s?rd=%s", host, string(ctx.XOriginalURL())), fasthttp.StatusUnauthorized)

			return
		}

		corsOrigin := ctx.Request.Header.Peek(headerOrigin)

		requestedWith := ctx.Request.Header.Peek("X-Requested-With")
		if requestedWith != nil {
			ctx.Logger.Warnf("Requested With: %v", requestedWith)
		}

		if corsOrigin != nil {
			corsOriginURL, err := url.Parse(string(corsOrigin))

			if err == nil && corsOriginURL != nil && utils.IsRedirectionSafe(*corsOriginURL, ctx.Configuration.Session.Domain) {
				ctx.Response.Header.SetBytesV(headerAccessControlAllowOrigin, corsOrigin)
				ctx.Response.Header.Set(headerVary, "Accept-Encoding, Origin")
				ctx.Response.Header.Set(headerAccessControlAllowCredentials, "false")

				corsHeaders := ctx.Request.Header.Peek(headerAccessControlRequestHeaders)
				if corsHeaders != nil {
					ctx.Response.Header.SetBytesV(headerAccessControlAllowHeaders, corsHeaders)
				}

				corsMethod := ctx.Request.Header.Peek(headerAccessControlRequestMethod)
				if corsHeaders != nil {
					ctx.Response.Header.SetBytesV(headerAccessControlAllowMethods, corsMethod)
				} else {
					ctx.Response.Header.Set(headerAccessControlAllowMethods, "GET")
				}
			}
		}

		next(ctx)
	}
}
