package xrpcuri

import (
	"github.com/reiver/go-xrpcuri/internal"
)

// ValidateID returns an error if the XRPC-URI 'id' is invalid.
// It returns nil if the XRPC-URI id is valid.
//
// NOTE THAT THIS IS NOT VALIDATING AN XRPC-URI, BUT IS INSTEAD VALIDATING AN XRPC-URI 'id'.
//
// An XRPC-URI 'id' must be an NSID, and its validation rules are defined at:
// https://atproto.com/specs/nsid
func ValidateID(id string) error {
	return xrpcuri_internal.ValidateID(id)
}
