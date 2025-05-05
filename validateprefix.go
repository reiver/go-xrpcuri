package xrpcuri

import (
	"github.com/reiver/go-erorr"

	"github.com/reiver/go-xrpcuri/enc"
	"github.com/reiver/go-xrpcuri/pln"
)

// ValidatePrefix only validates the prefix (i.e., "xrpc://" or "xrpc-unencrypted://") of a potential XRPC-URI or XRPC-unencrypted-URI.
//
// So, it checks to see if the URI starts with "xrpc://" or "xrpc-unencrypted://".
// And, that is it.
//
// You would use ValidatePrefix if you wanted to be very liberal in what you accept as a valid XRPC-URI or XRPC-unencrypted-URI, including accepting an empty 'authority' / 'host'.
// I.e., this does a bit more than [ValidatePrefix].
//
// ValidatePrefix calls [ValidateScheme] internally.
//
// For a more thorough validation of the whole XRPC-URI or XRPC-unencrypted-URI instead use [Validate].
func ValidatePrefix(uri string) error {

	err1 := xrpcuripln.ValidatePrefix(uri)
	err2 := xrpcurienc.ValidatePrefix(uri)

	switch {
	case nil == err1 || nil == err2:
		return nil
	default:
		const slashSlash string = "//"

		return erorr.Errorf("xrpcuri: URI %q is not an XRPC-URI or an XRPC-unencrypted-URI because it does not begin with %q or %q", uri, xrpcurienc.PrefixScheme + slashSlash, xrpcuripln.PrefixScheme + slashSlash)
	}
}
