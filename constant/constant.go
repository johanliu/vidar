package constant

import (
	"errors"
	"fmt"
	"net/http"
)

// MIME types
const (
	charsetUTF8 = "charset=UTF-8"
)

const (
	MIMEApplicationJSON                  = "application/json"
	MIMEApplicationJSONCharsetUTF8       = MIMEApplicationJSON + "; " + charsetUTF8
	MIMEApplicationJavaScript            = "application/javascript"
	MIMEApplicationJavaScriptCharsetUTF8 = MIMEApplicationJavaScript + "; " + charsetUTF8
	MIMEApplicationXML                   = "application/xml"
	MIMEApplicationXMLCharsetUTF8        = MIMEApplicationXML + "; " + charsetUTF8
	MIMETextXML                          = "text/xml"
	MIMETextXMLCharsetUTF8               = MIMETextXML + "; " + charsetUTF8
	MIMEApplicationForm                  = "application/x-www-form-urlencoded"
	MIMEApplicationProtobuf              = "application/protobuf"
	MIMEApplicationMsgpack               = "application/msgpack"
	MIMETextHTML                         = "text/html"
	MIMETextHTMLCharsetUTF8              = MIMETextHTML + "; " + charsetUTF8
	MIMETextPlain                        = "text/plain"
	MIMETextPlainCharsetUTF8             = MIMETextPlain + "; " + charsetUTF8
	MIMEMultipartForm                    = "multipart/form-data"
	MIMEOctetStream                      = "application/octet-stream"
)

// Header
const (
	HeaderAccept              = "Accept"
	HeaderAcceptEncoding      = "Accept-Encoding"
	HeaderAllow               = "Allow"
	HeaderAuthorization       = "Authorization"
	HeaderContentDisposition  = "Content-Disposition"
	HeaderContentEncoding     = "Content-Encoding"
	HeaderContentLength       = "Content-Length"
	HeaderContentType         = "Content-Type"
	HeaderCookie              = "Cookie"
	HeaderSetCookie           = "Set-Cookie"
	HeaderIfModifiedSince     = "If-Modified-Since"
	HeaderLastModified        = "Last-Modified"
	HeaderLocation            = "Location"
	HeaderUpgrade             = "Upgrade"
	HeaderVary                = "Vary"
	HeaderWWWAuthenticate     = "WWW-Authenticate"
	HeaderXForwardedFor       = "X-Forwarded-For"
	HeaderXForwardedProto     = "X-Forwarded-Proto"
	HeaderXForwardedProtocol  = "X-Forwarded-Protocol"
	HeaderXForwardedSsl       = "X-Forwarded-Ssl"
	HeaderXUrlScheme          = "X-Url-Scheme"
	HeaderXHTTPMethodOverride = "X-HTTP-Method-Override"
	HeaderXRealIP             = "X-Real-IP"
	HeaderXRequestID          = "X-Request-ID"
	HeaderServer              = "Server"
	HeaderOrigin              = "Origin"
)

// HTTP methods
const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
	HEAD   = "HEAD"
	PATCH  = "PATCH"
)

// Errors
var (
	BadRequestError           = NewHTTPError(http.StatusBadRequest)           //400
	UnauthorizedError         = NewHTTPError(http.StatusUnauthorized)         //401
	ForbiddenError            = NewHTTPError(http.StatusForbidden)            //403
	NotFoundError             = NewHTTPError(http.StatusNotFound)             //404
	MethodNotAllowedError     = NewHTTPError(http.StatusMethodNotAllowed)     //405
	UnsupportedMediaTypeError = NewHTTPError(http.StatusUnsupportedMediaType) //415
	ConflictError             = NewHTTPError(http.StatusConflict)             //409

	InternalServerError = NewHTTPError(http.StatusInternalServerError) //500
	NotImplementedError = NewHTTPError(http.StatusNotImplemented)      //501
)

var (
	TodoError   = errors.New("Todo Error")
	FormatError = errors.New("Format Error")
)

type HTTPError struct {
	Code    int
	Content string
}

func NewHTTPError(code int, message ...string) *HTTPError {
	var content string

	if len(message) > 0 {
		content = message[0]
	} else {
		content = http.StatusText(code)
	}

	return &HTTPError{Code: code, Content: content}
}

func (he *HTTPError) Error() string {
	return fmt.Sprintf("code=%d, message=%v", he.Code, he.Content)
}
