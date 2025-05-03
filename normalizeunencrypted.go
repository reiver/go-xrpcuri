package xrpcuri

import (
	"github.com/reiver/go-xrpcuri/pln"
)

// NormalizeUnencrypted returns the normalized form of an XRPC-unencrypted-URI.
//
// NormalizeUnencrypted does NOT validate the XRPC-unencrypted-URI.
// To validate, call [ValidateUnencrypted].
//
// An example of a non-normalized XRPC-unencrypted-URI would be:
//
//	xrpc-unencrypted://VIDEO.archive.ORG/COM.Example.fooBar
//
// Normalizing that non-normalized XRPC-URI would result in:
//
//	xrpc-unencrypted://video.archive.org/com.example.fooBar
func NormalizeUnencrypted(uri string) string {

	if err := ValidateSchemeUnencrypted(uri); nil != err {
		return uri
	}

	authority, id, query, fragment, err := xrpcuripln.Split(uri)
	if nil != err {
		return uri
	}

	authority = NormalizeAuthority(authority)
	id        = NormalizeID(id)

	return JoinUnencrypted(authority, id, query, fragment)
}
