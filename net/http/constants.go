package http

import (
	"net/http"
)

// This file contains references to constants (and "const" like variables) defined
// in net/http

const (
	MethodGet     = http.MethodGet
	MethodHead    = http.MethodHead
	MethodPost    = http.MethodPost
	MethodPut     = http.MethodPut
	MethodPatch   = http.MethodPatch
	MethodDelete  = http.MethodDelete
	MethodConnect = http.MethodConnect
	MethodOptions = http.MethodOptions
	MethodTrace   = http.MethodTrace
)

const (
	StatusContinue           = http.StatusContinue
	StatusSwitchingProtocols = http.StatusSwitchingProtocols
	StatusProcessing         = http.StatusProcessing
	StatusEarlyHints         = http.StatusEarlyHints

	StatusOK                   = http.StatusOK
	StatusCreated              = http.StatusCreated
	StatusAccepted             = http.StatusAccepted
	StatusNonAuthoritativeInfo = http.StatusNonAuthoritativeInfo
	StatusNoContent            = http.StatusNoContent
	StatusResetContent         = http.StatusResetContent
	StatusPartialContent       = http.StatusPartialContent
	StatusMultiStatus          = http.StatusMultiStatus
	StatusAlreadyReported      = http.StatusAlreadyReported
	StatusIMUsed               = http.StatusIMUsed

	StatusMultipleChoices  = http.StatusMultipleChoices
	StatusMovedPermanently = http.StatusMovedPermanently
	StatusFound            = http.StatusFound
	StatusSeeOther         = http.StatusSeeOther
	StatusNotModified      = http.StatusNotModified
	StatusUseProxy         = http.StatusUseProxy

	StatusTemporaryRedirect = http.StatusTemporaryRedirect
	StatusPermanentRedirect = http.StatusPermanentRedirect

	StatusBadRequest                   = http.StatusBadRequest
	StatusUnauthorized                 = http.StatusUnauthorized
	StatusPaymentRequired              = http.StatusPaymentRequired
	StatusForbidden                    = http.StatusForbidden
	StatusNotFound                     = http.StatusNotFound
	StatusMethodNotAllowed             = http.StatusMethodNotAllowed
	StatusNotAcceptable                = http.StatusNotAcceptable
	StatusProxyAuthRequired            = http.StatusProxyAuthRequired
	StatusRequestTimeout               = http.StatusRequestTimeout
	StatusConflict                     = http.StatusConflict
	StatusGone                         = http.StatusGone
	StatusLengthRequired               = http.StatusLengthRequired
	StatusPreconditionFailed           = http.StatusPreconditionFailed
	StatusRequestEntityTooLarge        = http.StatusRequestEntityTooLarge
	StatusRequestURITooLong            = http.StatusRequestURITooLong
	StatusUnsupportedMediaType         = http.StatusUnsupportedMediaType
	StatusRequestedRangeNotSatisfiable = http.StatusRequestedRangeNotSatisfiable
	StatusExpectationFailed            = http.StatusExpectationFailed
	StatusTeapot                       = http.StatusTeapot
	StatusMisdirectedRequest           = http.StatusMisdirectedRequest
	StatusUnprocessableEntity          = http.StatusUnprocessableEntity
	StatusLocked                       = http.StatusLocked
	StatusFailedDependency             = http.StatusFailedDependency
	StatusTooEarly                     = http.StatusTooEarly
	StatusUpgradeRequired              = http.StatusUpgradeRequired
	StatusPreconditionRequired         = http.StatusPreconditionRequired
	StatusTooManyRequests              = http.StatusTooManyRequests
	StatusRequestHeaderFieldsTooLarge  = http.StatusRequestHeaderFieldsTooLarge
	StatusUnavailableForLegalReasons   = http.StatusUnavailableForLegalReasons

	StatusInternalServerError           = http.StatusInternalServerError
	StatusNotImplemented                = http.StatusNotImplemented
	StatusBadGateway                    = http.StatusBadGateway
	StatusServiceUnavailable            = http.StatusServiceUnavailable
	StatusGatewayTimeout                = http.StatusGatewayTimeout
	StatusHTTPVersionNotSupported       = http.StatusHTTPVersionNotSupported
	StatusVariantAlsoNegotiates         = http.StatusVariantAlsoNegotiates
	StatusInsufficientStorage           = http.StatusInsufficientStorage
	StatusLoopDetected                  = http.StatusLoopDetected
	StatusNotExtended                   = http.StatusNotExtended
	StatusNetworkAuthenticationRequired = http.StatusNetworkAuthenticationRequired
)

const (
	DefaultMaxHeaderBytes      = http.DefaultMaxHeaderBytes
	DefaultMaxIdleConnsPerHost = http.DefaultMaxIdleConnsPerHost
	TimeFormat                 = http.TimeFormat
	TrailerPrefix              = http.TrailerPrefix
)

var (
	ErrNotSupported         = http.ErrNotSupported
	ErrUnexpectedTrailer    = http.ErrUnexpectedTrailer // Deprecated
	ErrMissingBoundary      = http.ErrMissingBoundary
	ErrNotMultipart         = http.ErrNotMultipart
	ErrHeaderTooLong        = http.ErrHeaderTooLong        // Deprecated
	ErrShortBody            = http.ErrShortBody            // Deprecated
	ErrMissingContentLength = http.ErrMissingContentLength // Deprecated

	ErrBodyNotAllowed  = http.ErrBodyNotAllowed
	ErrHijacked        = http.ErrHijacked
	ErrContentLength   = http.ErrContentLength
	ErrWriteAfterFlush = http.ErrWriteAfterFlush // Deprecated

	ServerContextKey    = http.ServerContextKey
	LocalAddrContextKey = http.LocalAddrContextKey

	DefaultClient         = http.DefaultClient
	DefaultServeMux       = http.DefaultServeMux
	ErrAbortHandler       = http.ErrAbortHandler
	ErrBodyReadAfterClose = http.ErrBodyReadAfterClose
	ErrHandlerTimeout     = http.ErrHandlerTimeout
	ErrLineTooLong        = http.ErrLineTooLong
	ErrMissingFile        = http.ErrMissingFile
	ErrNoCookie           = http.ErrNoCookie
	ErrNoLocation         = http.ErrNoLocation
	ErrSchemeMismatch     = http.ErrSchemeMismatch
	ErrServerClosed       = http.ErrServerClosed
	ErrSkipAltProtocol    = http.ErrSkipAltProtocol
	ErrUseLastResponse    = http.ErrUseLastResponse
	NoBody                = http.NoBody
)
