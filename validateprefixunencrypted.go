package xrpcuri

import (
	"github.com/reiver/go-erorr"

	"github.com/reiver/go-xrpcuri/internal"
)

// ValidatePrefixUnencrypted only validates the prefix (i.e., "xrpc://") of a potential XRPC-unencrypted-URI.
//
// So, it checks to see if the URI starts with "xrpc-unencrypted://".
// And, that is it.
//
// You would use ValidatePrefixUnencrypted if you wanted to be very liberal in what you accept as a valid XRPC-unencrypted-URI, including accepting an empty 'host'.
// I.e., this does a bit more than [ValidatePrefixUnencrypted].
//
// ValidatePrefixUnencrypted calls [ValidateSchemeUnencrypted] internally.
//
// For a more thorough validation of the whole XRPC-unencrypted-URI instead use [ValidateUnencrypted].
func ValidatePrefixUnencrypted(uri string) error {

	if err := ValidateSchemeUnencrypted(uri); nil != err {
		return err
	}

	str := uri[xrpcuri_internal.LenPrefixSchemeUnencrypted:]

	{
		const slashSlash string = "//"
		const lenSlashSlash int = len(slashSlash)

		var lenstr int = len(str)
		if lenstr < lenSlashSlash {
			return erorr.Errorf("xrpcuri: XRPC-unencrypted-URI %q is not valid because it does not have %q after \"xrpc-unencrypted:\" — too short", uri, slashSlash)
		}

		var beginning string = str[:lenSlashSlash]

		if beginning != slashSlash {
			return erorr.Errorf("xrpcuri: XRPC-unencrypted-URI %q is not valid because it does not have %q after \"xrpc-unencrypted:\"", uri, slashSlash)
		}
	}

	return nil
}
