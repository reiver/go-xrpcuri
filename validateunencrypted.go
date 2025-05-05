package xrpcuri

import (
	"github.com/reiver/go-xrpcuri/internal"
	"github.com/reiver/go-xrpcuri/pln"
)

// ValidateUnencrypted returns an error if the XRPC-URI is invalid.
// It returns nil if the XRPC-URI is valid.
//
// To validate a potential XRPC-URI more liberally, use one of the following:
// [ValidatePrefixUnencrypted],
// [ValidateSchemeUnencrypted].
func ValidateUnencrypted(uri string) error {

	if err := xrpcuripln.ValidatePrefix(uri); nil != err {
		return err
	}

	authority, id, _, _, err := xrpcuripln.Split(uri)
	if nil != err {
		return err
	}

	if err := xrpcuri_internal.ValidateAuthorityPretty(authority, uri); nil != err {
		 return err
	}
	if err := xrpcuri_internal.ValidateIDPretty(id, uri); nil != err {
		 return err
	}

	return nil
}
