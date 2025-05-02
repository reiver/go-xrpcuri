package xrpcuri

import (
	"strings"

	"github.com/reiver/go-erorr"
)

const (
	errEmptyAuthority    = erorr.Error("xrpcuri: empty authority")
	errAtSignInAuthority = erorr.Error("xrpcuri: authority may not have an \"@\" in it")
)

// ValidateAuthority returns an error if the XRPC-URI 'authority' is invalid.
// It returns nil if the XRPC-URI authority is valid.
//
// NOTE THAT THIS IS NOT VALIDATING AN XRPC-URI, BUT IS INSTEAD VALIDATING AN XRPC-URI 'authority'.
func ValidateAuthority(authority string) error {

	if "" == authority {
		return errEmptyAuthority
	}

	{
		const disallowed string = "@"

		if strings.Contains(authority, disallowed) {
			return errAtSignInAuthority
		}
	}

	return nil
}

func validateAuthority(authority string, uri string) error {

	if err := ValidateAuthority(authority); nil != err {
		switch {
		case erorr.Is(err, errEmptyAuthority):
			return erorr.Errorf("xrpcuri: XRPC-URI %q has an empty 'authority'", uri)
		case erorr.Is(err, errAtSignInAuthority):
			return erorr.Errorf("xrpcuri: XRPC-URI %q may not have an \"@\" in its authority %q", uri, authority)
		default:
			return erorr.Errorf("xrpcuri: XRPC-URI %q has a authority %q that is not a valid: %w", uri, authority, err)
		}
	}

	return nil
}
