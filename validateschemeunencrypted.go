package xrpcuri

import (
	"strings"

	"github.com/reiver/go-erorr"

	"github.com/reiver/go-xrpcuri/internal"
)

// ValidateSchemeUnencrypted only validates the scheme of a potential XRPC-unencrypted-URI.
//
// So, it checks to see if the URI starts with "xrpc-unencrypted:".
// And, that is it.
//
// You would use ValidateSchemeUnencrypted if you wanted to be very liberal in what you accept as a valid XRPC-unencrypted-URI, including not caring if the XRPC-unencrypted-URI has a 'host' or not.
// I.e., this is the minimum amount of validation you can do to validate an XRPC-unencrypted-URI.
//
// Note that you are passing ValidateSchemeUnencrypted the whole URI, and not just the scheme.
//
// For a more thorough validation of the whole XRPC-unencrypted-URI instead use [ValidateUnencrypted].
//
// Alternatively, to validate a bit more than ValidateSchemeUnencrypted, without being as thorough as [ValidateUnencrypted], instead use [ValidatePrefixUnencrypted].
func ValidateSchemeUnencrypted(uri string) error {

	if "" == uri {
		return xrpcuri_internal.ErrEmptyURI
	}

	{
		var lenuri int = len(uri)
		if lenuri < xrpcuri_internal.LenPrefixSchemeUnencrypted {
			return erorr.Errorf("xrpcuri: URI %q is not an XRPC-unencrypted-URI because it does not begin with %q", uri, xrpcuri_internal.PrefixSchemeUnencrypted)
		}

                var beginning string = uri[:xrpcuri_internal.LenPrefixSchemeUnencrypted]

                if strings.ToLower(beginning) != xrpcuri_internal.PrefixSchemeUnencrypted {
			return erorr.Errorf("xrpcuri: URI %q is not an XRPC-unencrypted-URI because it does not begin with %q", uri, xrpcuri_internal.PrefixSchemeUnencrypted)
                }
	}

	return nil
}
