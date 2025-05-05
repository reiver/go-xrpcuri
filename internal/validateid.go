package xrpcuri_internal

import (
	"strings"

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

func ValidateIDPretty(id string, uri string) error {
	var kind string = "XRPC-URI"
	if strings.HasPrefix(uri, PrefixSchemeUnencrypted) {
		kind = "XRPC-unencrypted-URI"
	}


	if err := ValidateID(id); nil != err {
		return erorr.Errorf("xrpcuri: %s %q has an id %q that is not a valid NSID: %w", kind, uri, id, err)
	}

	return nil
}
