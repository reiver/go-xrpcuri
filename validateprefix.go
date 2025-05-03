package xrpcuri

import (
	"github.com/reiver/go-erorr"

	"github.com/reiver/go-xrpcuri/internal"
)

// ValidatePrefix only validates the prefix (i.e., "xrpc://") of a potential XRPC-URI.
//
// So, it checks to see if the URI starts with "xrpc://".
// And, that is it.
//
// You would use ValidatePrefix if you wanted to be very liberal in what you accept as a valid XRPC-URI, including accepting an empty 'host'.
// I.e., this does a bit more than [ValidatePrefix].
//
// ValidatePrefix calls [ValidateScheme] internally.
//
// For a more thorough validation of the whole XRPC-URI instead use [Validate].
func ValidatePrefix(uri string) error {

	if err := ValidateScheme(uri); nil != err {
		return err
	}

	str := uri[xrpcuri_internal.LenPrefixScheme:]

	{
		const slashSlash string = "//"
		const lenSlashSlash int = len(slashSlash)

		var lenstr int = len(str)
		if lenstr < lenSlashSlash {
			return erorr.Errorf("xrpcuri: XRPC-URI %q is not valid because it does not have %q after \"xrpc:\" — too short", uri, slashSlash)
		}

		var beginning string = str[:lenSlashSlash]

		if beginning != slashSlash {
			return erorr.Errorf("xrpcuri: XRPC-URI %q is not valid because it does not have %q after \"xrpc:\"", uri, slashSlash)
		}
	}

	return nil
}
