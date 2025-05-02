package xrpcuri

import (
	"github.com/reiver/go-erorr"
	"github.com/reiver/go-nsid"
)

// ValidateID returns an error if the XRPC-URI 'id' is invalid.
// It returns nil if the XRPC-URI id is valid.
//
// NOTE THAT THIS IS NOT VALIDATING AN XRPC-URI, BUT IS INSTEAD VALIDATING AN XRPC-URI 'id'.
//
// An XRPC-URI 'id' must be an NSID, and its validation rules are defined at:
// https://atproto.com/specs/nsid
func ValidateID(id string) error {
	return nsid.Validate(id)
}

func validateID(id string, uri string) error {
	if err := ValidateID(id); nil != err {
		return erorr.Errorf("xrpcuri: XRPC-URI %q has a id %q that is not a valid NSID: %w", uri, id, err)
	}

	return nil
}
