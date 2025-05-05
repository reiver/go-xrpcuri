package xrpcuri

import (
	"github.com/reiver/go-xrpcuri/internal"
)

// SplitAndNormalizeAndValidate returns the 'authority', 'id', 'query', and 'fragment' of an XRPC-URI or XRPC-unencrypted-URI.
//
// An 'id' is an NSID (Namespaced Identifier).
//
// For example:
//
//	var uri string = "at://Host.EXAMPLE/COM.Example.foorBar"
//	
//	authority, id, query, fragment, err := xrpcuri.SplitAndNormalizeAndValidate(uri)
//	if nil != err {
//		return err
//	}
//	
//	// authority == "host.example"
//	// id        == "com.example.foorBar"
//	// query     == ""
//	// fragment  == ""
//
// SplitAndNormalizeAndValidate normalizes the returned values.
//
// If you are not sure whether to use [Split] or [SpliAndNormalize] or SplitAndNormalizeAndValidate, use [SplitAndNormalizeAndValidate].
func SplitAndNormalizeAndValidate(uri string) (authority string, id string, query string, fragment string, err error) {
	if "" == uri {
                err = xrpcuri_internal.ErrEmptyURI
		return
	}

	authority, id, query, fragment, err = SplitAndNormalize(uri)
	if nil != err {
		return
	}

	{
		if err = xrpcuri_internal.ValidateAuthorityPretty(authority, uri); nil != err {
			 return
		}
		if err = xrpcuri_internal.ValidateIDPretty(id, uri); nil != err {
			return
		}
	}

	return
}
