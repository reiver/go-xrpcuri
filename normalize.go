package xrpcuri

import (
	"github.com/reiver/go-xrpcuri/enc"
	"github.com/reiver/go-xrpcuri/pln"
)

// Normalize returns the normalized form of an XRPC-URI or an XRPC-unencrypted-URI.
//
// Normalize does NOT validate the XRPC-URI or XRPC-unencrypted-URI.
// To validate, call [Validate].
//
// An example of a non-normalized XRPC-URI would be:
//
//	xrpc://VIDEO.archive.ORG/COM.Example.fooBar
//
// Normalizing that non-normalized XRPC-URI would result in:
//
//	xrpc://video.archive.org/com.example.fooBar
func Normalize(uri string) string {
	switch {
	case nil == xrpcuripln.ValidateScheme(uri):
		return xrpcuripln.Normalize(uri)
	default:
		return xrpcurienc.Normalize(uri)
	}
}
