package xrpcuri

import (
	"github.com/reiver/go-xrpcuri/enc"
	"github.com/reiver/go-xrpcuri/internal"
	"github.com/reiver/go-xrpcuri/pln"
)

// Validate returns an error if the XRPC-URI or XRPC-unencrypted-URI is invalid.
// It returns nil if the XRPC-URI is valid.
//
// To validate a potential XRPC-URI or XRPC-unencrypted-URI more liberally, use one of the following:
// [ValidatePrefix],
// [ValidateScheme].
func Validate(uri string) error {

	if "" == uri {
		return xrpcuri_internal.ErrEmptyURI
	}

	err1 := xrpcuripln.Validate(uri)
	err2 := xrpcurienc.Validate(uri)

	switch {
	case nil == err1 || nil == err2:
		return nil
	default:
		return err2
	}
}
