package xrpcuri

import (
	"github.com/reiver/go-xrpcuri/enc"
	"github.com/reiver/go-xrpcuri/pln"
)

// Split returns the 'host', 'id', 'query', and 'fragment' of at XRPC-URI or a XRPC-unencrypted-URI.
//
// The 'id' should be an NSID (Namespaced Identifier).
//
// For example:
//
//	var uri string = "xrpc://public.api.bsky.app/app.bsky.actor.getProfile?actor=reiver.bsky.social"
//
//	host, id, query, fragment, err := xrpcuri.Split(uri)
//	if nil != err {
//		return err
//	}
//
//	// host     == "public.api.bsky.app"
//	// id       == "app.bsky.actor.getProfile"
//	// query    == "actor=reiver.bsky.social"
//	// fragment == ""
//
// Split does NOT normalize the returned values.
//
// If you are not sure whether to use Split or [SpliAndNormalize], use [SplitAndNormalize].
func Split(uri string) (host string, id string, query string, fragment string, err error) {
	switch {
	case nil == xrpcuripln.ValidateScheme(uri):
		return xrpcuripln.Split(uri)
	default:
		return xrpcurienc.Split(uri)
	}
}
