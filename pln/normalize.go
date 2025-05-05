package xrpcuripln

import (
	"github.com/reiver/go-xrpcuri/internal"
)

// Normalize returns the normalized form of an XRPC-unencrypted-URI.
//
// Normalize does NOT validate the XRPC-unencrypted-URI.
// To validate, call [Validate].
//
// An example of a non-normalized XRPC-unencrypted-URI would be:
//
//	xrpc-unencrypted://VIDEO.archive.ORG/COM.Example.fooBar
//
// Normalizing that non-normalized XRPC-unencrypted-URI would result in:
//
//	xrpc-unencrypted://video.archive.org/com.example.fooBar
func Normalize(uri string) string {

	if err := ValidateScheme(uri); nil != err {
		return uri
	}

	authority, id, query, fragment, err := Split(uri)
	if nil != err {
		return uri
	}

	authority = xrpcuri_internal.NormalizeAuthority(authority)
	id        = xrpcuri_internal.NormalizeID(id)

	return Join(authority, id, query, fragment)
}
