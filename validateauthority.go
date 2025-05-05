package xrpcuri

import (
	"github.com/reiver/go-xrpcuri/internal"
)

// ValidateAuthority returns an error if the XRPC-URI 'authority' is invalid.
// It returns nil if the XRPC-URI authority is valid.
//
// NOTE THAT THIS IS NOT VALIDATING AN XRPC-URI, BUT IS INSTEAD VALIDATING AN XRPC-URI 'authority'.
func ValidateAuthority(authority string) error {
	return xrpcuri_internal.ValidateAuthority(authority)
}
