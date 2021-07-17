package middlewares

import "github.com/valyala/fasthttp"

func AutomaticCORSMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		corsOrigin := ctx.Request.Header.Peek("Origin")

		// TODO: Check Origin is protected by Authelia.
		if corsOrigin != nil {
			ctx.Response.Header.SetBytesV("Access-Control-Allow-Origin", corsOrigin)

			ctx.Response.Header.SetBytesV("Access-Control-Allow-Headers", ctx.Request.Header.Peek("Access-Control-Request-Headers"))
			ctx.Response.Header.SetBytesV("Access-Control-Allow-Methods", ctx.Request.Header.Peek("Access-Control-Request-Method"))
		}

		next(ctx)
	}
}
