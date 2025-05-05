package xrpcuripln

import (
	"strings"

	"github.com/reiver/go-erorr"

	"github.com/reiver/go-xrpcuri/internal"
)

// ValidateScheme only validates the scheme of a potential XRPC-unencrypted-URI.
//
// So, it checks to see if the URI starts with "xrpc-unencrypted:".
// And, that is it.
//
// You would use ValidateScheme if you wanted to be very liberal in what you accept as a valid XRPC-unencrypted-URI, including not caring if the XRPC-unencrypted-URI has a 'host' or not.
// I.e., this is the minimum amount of validation you can do to validate an XRPC-unencrypted-URI.
//
// Note that you are passing ValidateScheme the whole URI, and not just the scheme.
//
// For a more thorough validation of the whole XRPC-unencrypted-URI instead use [Validate].
//
// Alternatively, to validate a bit more than ValidateScheme, without being as thorough as [Validate], instead use [ValidatePrefix].
func ValidateScheme(uri string) error {

	if "" == uri {
		return xrpcuri_internal.ErrEmptyURI
	}

	{
		var lenuri int = len(uri)
		if lenuri < LenPrefixScheme {
			return erorr.Errorf("xrpcuri: URI %q is not an XRPC-unencrypted-URI because it does not begin with %q", uri, PrefixScheme)
		}

                var beginning string = uri[:LenPrefixScheme]

                if strings.ToLower(beginning) != PrefixScheme {
			return erorr.Errorf("xrpcuri: URI %q is not an XRPC-unencrypted-URI because it does not begin with %q", uri, PrefixScheme)
                }
	}

	return nil
}
