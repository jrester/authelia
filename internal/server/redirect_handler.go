package server

import (
	"github.com/valyala/fasthttp"

	"github.com/authelia/authelia/internal/middlewares"
)

func handleRedirect(ctx *middlewares.AutheliaCtx) {
	ctx.Redirect("/?"+ctx.QueryArgs().String(), fasthttp.StatusMovedPermanently)
}
