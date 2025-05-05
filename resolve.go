package xrpcuri

import (
	"github.com/reiver/go-erorr"

	"github.com/reiver/go-xrpcuri/enc"
	"github.com/reiver/go-xrpcuri/internal"
	"github.com/reiver/go-xrpcuri/pln"
)

// Resolve turns an XRPC-URI into an HTTP-URI.
//
// For example:
//
//	"xrpc://public.api.bsky.app/app.bsky.actor.getProfile" -> "https://public.api.bsky.app/xrpc/app.bsky.actor.getProfile"
//
//	"xrpc://example.com/com.example.fooBar" -> "https://example.com/xrpc/com.example.fooBar"
//
//	"xrpc-unencrypted://example.com/com.example.fooBar" -> "http://example.com/xrpc/com.example.fooBar"
func Resolve(uri string, requestType string) (string, error) {
	switch requestType {
	case RequestTypeExecute, RequestTypeProcedure:
		return resolveWeb(uri)
	case RequestTypeQuery:
		return resolveWeb(uri)
	case RequestTypeSubscribe:
		return resolveWebSocket(uri)
	default:
		const nada string = ""
		return nada, erorr.Errorf("xrpcuri: do not know how to resolve a %q XRPC request-type", requestType)
	}
}

func resolveWeb(uri string) (string, error) {
	var fn     func(string)(string,string,string,string,error)
	var scheme string

	switch {
	case nil == xrpcuripln.ValidateScheme(uri):
		fn     = xrpcuripln.Split
		scheme = "http"
	default:
		fn     = xrpcurienc.Split
		scheme = "https"
	}

	return xrpcuri_internal.Resolve(uri, fn, scheme)
}

func resolveWebSocket(uri string) (string, error) {
	var fn     func(string)(string,string,string,string,error)
	var scheme string

	switch {
	case nil == xrpcuripln.ValidateScheme(uri):
		fn     = xrpcuripln.Split
		scheme = "ws"
	default:
		fn     = xrpcurienc.Split
		scheme = "wss"
	}

	return xrpcuri_internal.Resolve(uri, fn, scheme)
}
