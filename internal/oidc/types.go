package oidc

import (
	"crypto/rsa"

	"github.com/ory/fosite"
	"github.com/ory/fosite/handler/openid"
	"github.com/ory/fosite/storage"
	"github.com/ory/herodot"
	"gopkg.in/square/go-jose.v2"

	"github.com/authelia/authelia/internal/authorization"
)

// OpenIDConnectProvider for OpenID Connect.
type OpenIDConnectProvider struct {
	Fosite     fosite.OAuth2Provider
	Store      *OpenIDConnectStore
	KeyManager *KeyManager

	herodot *herodot.JSONWriter
}

// OpenIDConnectStore is Authelia's internal representation of the fosite.Storage interface.
//
//	Currently it is mostly just implementing a decorator pattern other then GetInternalClient.
//	The long term plan is to have these methods interact with the Authelia storage and
//	session providers where applicable.
type OpenIDConnectStore struct {
	clients map[string]*InternalClient
	memory  *storage.MemoryStore
}

// InternalClient represents the client internally.
type InternalClient struct {
	ID            string   `json:"id"`
	Description   string   `json:"-"`
	Secret        []byte   `json:"client_secret,omitempty"`
	RedirectURIs  []string `json:"redirect_uris"`
	GrantTypes    []string `json:"grant_types"`
	ResponseTypes []string `json:"response_types"`
	Scopes        []string `json:"scopes"`
	Audience      []string `json:"audience"`
	Public        bool     `json:"public"`

	ResponseModes []fosite.ResponseModeType `json:"response_modes"`

	UserinfoSigningAlgorithm string `json:"userinfo_signed_response_alg,omitempty"`

	Policy authorization.Level `json:"-"`
}

// KeyManager keeps track of all of the active/inactive rsa keys and provides them to services requiring them.
// It additionally allows us to add keys for the purpose of key rotation in the future.
type KeyManager struct {
	activeKeyID string
	keys        map[string]*rsa.PrivateKey
	keySet      *jose.JSONWebKeySet
	strategy    *RS256JWTStrategy
}

// AutheliaHasher implements the fosite.Hasher interface without an actual hashing algo.
type AutheliaHasher struct{}

// ConsentGetResponseBody schema of the response body of the consent GET endpoint.
type ConsentGetResponseBody struct {
	ClientID          string     `json:"client_id"`
	ClientDescription string     `json:"client_description"`
	Scopes            []Scope    `json:"scopes"`
	Audience          []Audience `json:"audience"`
}

// Scope represents the scope information.
type Scope struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Audience represents the audience information.
type Audience struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// WellKnownConfiguration is the OIDC well known config struct. The following spec's document this struct:
//
// OpenID Connect Discovery https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata
//
// OpenID Connect Session Management https://openid.net/specs/openid-connect-session-1_0-17.html#OPMetadata
//
// OAuth 2.0 Authorization Server Metadata https://datatracker.ietf.org/doc/html/rfc8414#section-2
type WellKnownConfiguration struct {
	Issuer  string `json:"issuer"`
	JWKSURI string `json:"jwks_uri"`

	AuthorizationEndpoint string `json:"authorization_endpoint"`
	TokenEndpoint         string `json:"token_endpoint"`
	IntrospectionEndpoint string `json:"introspection_endpoint,omitempty"`
	UserinfoEndpoint      string `json:"userinfo_endpoint,omitempty"`
	RevocationEndpoint    string `json:"revocation_endpoint,omitempty"`
	RegistrationEndpoint  string `json:"registration_endpoint,omitempty"`

	TokenEndpointAuthSigningAlgValuesSupported         []string `json:"token_endpoint_auth_signing_alg_values_supported,omitempty"`
	IntrospectionEndpointAuthSigningAlgValuesSupported []string `json:"introspection_endpoint_auth_signing_alg_values_supported,omitempty"`
	IDTokenSigningAlgValuesSupported                   []string `json:"id_token_signing_alg_values_supported"`
	UserinfoSigningAlgValuesSupported                  []string `json:"userinfo_signing_alg_values_supported,omitempty"`
	RequestObjectSigningAlgValuesSupported             []string `json:"request_object_signing_alg_values_supported,omitempty"`
	CodeChallengeMethodsSupported                      []string `json:"code_challenge_methods_supported,omitempty"`

	IDTokenEncryptionAlgValuesSupported       []string `json:"id_token_encryption_alg_values_supported,omitempty"`
	UserinfoEncryptionAlgValuesSupported      []string `json:"userinfo_encryption_alg_values_supported,omitempty"`
	RequestObjectEncryptionAlgValuesSupported []string `json:"request_object_encryption_alg_values_supported,omitempty"`

	IDTokenEncryptionEncValuesSupported       []string `json:"id_token_encryption_enc_values_supported,omitempty"`
	UserinfoEncryptionEncValuesSupported      []string `json:"userinfo_encryption_enc_values_supported,omitempty"`
	RequestObjectEncryptionEncValuesSupported []string `json:"request_object_encryption_enc_values_supported,omitempty"`

	TokenEndpointAuthMethodsSupported         []string `json:"token_endpoint_auth_methods_supported,omitempty"`
	IntrospectionEndpointAuthMethodsSupported []string `json:"introspection_endpoint_auth_methods_supported,omitempty"`
	RevocationEndpointAuthMethodsSupported    []string `json:"revocation_endpoint_auth_methods_supported,omitempty"`

	ACRValuesSupported     []string `json:"acr_values_supported,omitempty"`
	DisplayValuesSupported []string `json:"display_values_supported,omitempty"`
	SubjectTypesSupported  []string `json:"subject_types_supported"`
	ResponseModesSupported []string `json:"response_modes_supported,omitempty"`
	ResponseTypesSupported []string `json:"response_types_supported"`
	GrantTypesSupported    []string `json:"grant_types_supported,omitempty"`
	ScopesSupported        []string `json:"scopes_supported,omitempty"`
	ClaimsSupported        []string `json:"claims_supported,omitempty"`
	ClaimTypesSupported    []string `json:"claim_types_supported,omitempty"`
	UILocalesSupported     []string `json:"ui_locales_supported,omitempty"`

	ClaimsParameterSupported           bool `json:"claims_parameter_supported"`
	RequestParameterSupported          bool `json:"request_parameter_supported"`
	RequestURIParameterSupported       bool `json:"request_uri_parameter_supported"`
	RequireRequestURIRegistration      bool `json:"require_request_uri_registration"`
	BackChannelLogoutSupported         bool `json:"backchannel_logout_supported"`
	FrontChannelLogoutSupported        bool `json:"frontchannel_logout_supported"`
	BackChannelLogoutSessionSupported  bool `json:"backchannel_logout_session_supported"`
	FrontChannelLogoutSessionSupported bool `json:"frontchannel_logout_session_supported"`

	ServiceDocumentation string `json:"service_documentation,omitempty"`
	OPPolicyURI          string `json:"op_policy_uri,omitempty"`
	OPTOSURI             string `json:"op_tos_uri,omitempty"`
}

// OpenIDSession holds OIDC Session information.
type OpenIDSession struct {
	*openid.DefaultSession `json:"idToken"`

	Extra    map[string]interface{} `json:"extra"`
	ClientID string
}
