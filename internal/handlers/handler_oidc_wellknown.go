package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"

	"github.com/authelia/authelia/internal/middlewares"
	"github.com/authelia/authelia/internal/oidc"
)

func oidcWellKnown(ctx *middlewares.AutheliaCtx) {
	// TODO (james-d-elliott): append the server.path here for path based installs. Also check other instances in OIDC.
	issuer, err := ctx.ForwardedProtoHost()
	if err != nil {
		ctx.Logger.Errorf("Error occurred in ForwardedProtoHost: %+v", err)
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)

		return
	}

	wellKnown := oidc.WellKnownConfiguration{
		Issuer:  issuer,
		JWKSURI: fmt.Sprintf("%s%s", issuer, oidcJWKsPath),

		AuthorizationEndpoint: fmt.Sprintf("%s%s", issuer, oidcAuthorizationPath),
		TokenEndpoint:         fmt.Sprintf("%s%s", issuer, oidcTokenPath),
		IntrospectionEndpoint: fmt.Sprintf("%s%s", issuer, oidcIntrospectionPath),
		UserinfoEndpoint:      fmt.Sprintf("%s%s", issuer, oidcUserinfoPath),
		RevocationEndpoint:    fmt.Sprintf("%s%s", issuer, oidcRevocationPath),

		TokenEndpointAuthSigningAlgValuesSupported: []string{"RS256"},
		IDTokenSigningAlgValuesSupported:           []string{"RS256"},
		UserinfoSigningAlgValuesSupported:          []string{"none", "RS256"},
		RequestObjectSigningAlgValuesSupported:     []string{"none", "RS256"},
		CodeChallengeMethodsSupported:              []string{"plain", "S256"},

		TokenEndpointAuthMethodsSupported:         []string{"client_secret_post", "client_secret_basic", "private_key_jwt", "none"},
		IntrospectionEndpointAuthMethodsSupported: []string{"client_secret_post", "client_secret_basic", "private_key_jwt", "none"},
		RevocationEndpointAuthMethodsSupported:    []string{"client_secret_post", "client_secret_basic", "private_key_jwt", "none"},

		DisplayValuesSupported: []string{
			"page",
		},
		SubjectTypesSupported: []string{
			"public",
		},
		ResponseModesSupported: []string{
			"form_post",
			"query",
			"fragment",
		},
		ResponseTypesSupported: []string{
			"code",
			"token",
			"id_token",
			"code token",
			"code id_token",
			"token id_token",
			"code token id_token",
			"none",
		},
		GrantTypesSupported: []string{
			"authorization_code",
			"implicit",
			"client_credentials",
			"refresh_token",
		},
		ScopesSupported: []string{
			"openid",
			"offline_access",
			"profile",
			"groups",
			"email",
		},
		ClaimsSupported: []string{
			"aud",
			"exp",
			"iat",
			"iss",
			"jti",
			"rat",
			"sub",
			"auth_time",
			"nonce",
			"email",
			"email_verified",
			"alt_emails",
			"groups",
			"name",
		},
		ClaimTypesSupported: []string{
			"normal",
		},
		UILocalesSupported: []string{
			"en-US",
		},

		ClaimsParameterSupported:           true,
		RequestParameterSupported:          false,
		RequestURIParameterSupported:       false,
		BackChannelLogoutSupported:         false,
		FrontChannelLogoutSupported:        false,
		BackChannelLogoutSessionSupported:  false,
		FrontChannelLogoutSessionSupported: false,
	}

	ctx.SetContentType("application/json")

	if err := json.NewEncoder(ctx).Encode(wellKnown); err != nil {
		ctx.Logger.Errorf("Error occurred in json Encode: %+v", err)
		// TODO: Determine if this is the appropriate error code here.
		ctx.Response.SetStatusCode(fasthttp.StatusInternalServerError)

		return
	}
}
