package xrpcuri

import (
	"strings"

	"github.com/reiver/go-erorr"
)

// ValidateScheme only validates the scheme of a potential XRPC-URI.
//
// So, it checks to see if the URI starts with "xrpc:".
// And, that is it.
//
// You would use ValidateScheme if you wanted to be very liberal in what you accept as a valid XRPC-URI, including not caring if the XRPC-URI has a 'host' or not.
// I.e., this is the minimum amount of validation you can do to validate an XRPC-URI.
//
// Note that you are passing ValidateScheme the whole URI, and not just the scheme.
//
// For a more thorough validation of the whole XRPC-URI instead use [Validate].
//
// Alternatively, to validate a bit more than ValidateScheme, without being as thorough as [Validate], instead use [ValidatePrefix].
func ValidateScheme(uri string) error {

	if "" == uri {
		return errEmptyURI
	}

	{
		var lenuri int = len(uri)
		if lenuri < lenPrefixScheme {
			return erorr.Errorf("xrpcuri: URI %q is not an XRPC-URI because it does not begin with %q", uri, prefixScheme)
		}

                var beginning string = uri[:lenPrefixScheme]

                if strings.ToLower(beginning) != prefixScheme {
			return erorr.Errorf("xrpcuri: URI %q is not an XRPC-URI because it does not begin with %q", uri, prefixScheme)
                }
	}

	return nil
}
