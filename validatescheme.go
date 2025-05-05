package xrpcuri

import (
	"github.com/reiver/go-erorr"

	"github.com/reiver/go-xrpcuri/enc"
	"github.com/reiver/go-xrpcuri/pln"
)

// ValidateScheme only validates the scheme of a potential XRPC-URI or XRPC-unencrypted-URI.
//
// So, it checks to see if the URI starts with "xrpc:" or "xrpc-unencrypted:".
// And, that is it.
//
// You would use ValidateScheme if you wanted to be very liberal in what you accept as a valid XRPC-URI or XRPC-unencrypted-URI, including not caring if the XRPC-URI or XRPC-unencrypted-URI has a 'authority' / 'host' or not.
// I.e., this is the minimum amount of validation you can do to validate an XRPC-URI or an XRPC-unencrypted-URI.
//
// Note that you are passing ValidateScheme the whole URI, and not just the scheme.
//
// For a more thorough validation of the whole XRPC-URI pr XRPC-unencrypted-URI instead use [Validate].
//
// Alternatively, to validate a bit more than ValidateScheme, without being as thorough as [Validate], instead use [ValidatePrefix].
func ValidateScheme(uri string) error {

	err1 := xrpcuripln.ValidateScheme(uri)
	err2 := xrpcurienc.ValidateScheme(uri)

	switch {
	case nil == err1 || nil == err2:
		return nil
	default:
		return erorr.Errorf("xrpcuri: URI %q is not an XRPC-URI or an XRPC-unencrypted-URI because it does not begin with %q or %q", uri, xrpcurienc.PrefixScheme, xrpcuripln.PrefixScheme)
	}
}
