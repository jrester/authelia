package middlewares

const jwtIssuer = "Authelia"

var (
	contentTypeApplicationJSON = []byte("application/json")
	contentTypeTextHTML        = []byte("text/html")

	headerAccept           = []byte("Accept")
	headerLocation         = []byte("Location")
	headerXForwardedProto  = []byte("X-Forwarded-Proto")
	headerXForwardedMethod = []byte("X-Forwarded-Method")
	headerXForwardedHost   = []byte("X-Forwarded-Host")
	headerXForwardedFor    = []byte("X-Forwarded-For")
	headerXForwardedURI    = []byte("X-Forwarded-URI")
	headerXOriginalURL     = []byte("X-Original-URL")
	headerXRequestedWith   = []byte("X-Requested-With")

	headerValueXRequestedWithXHR = []byte("XMLHttpRequest")
)

var okMessageBytes = []byte("{\"status\":\"OK\"}")

const (
	messageOperationFailed                      = "Operation failed"
	messageIdentityVerificationTokenAlreadyUsed = "The identity verification token has already been used"
	messageIdentityVerificationTokenHasExpired  = "The identity verification token has expired"
)

var protoHostSeparator = []byte("://")
