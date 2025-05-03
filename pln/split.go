package xrpcuripln

import (
	"github.com/reiver/go-xrpcuri/internal"
)

// Split returns the 'host', 'id', 'query', and 'fragment' of at XRPC-unencrypted-URI.
//
// A 'id' should be an NSID (Namespaced Identifier).
//
// For example:
//
//	var uri string = "xrpc-unencrypted://public.api.bsky.app/app.bsky.actor.getProfile?actor=reiver.bsky.social"
//
//	host, id, query, fragment, err := xrpcuripln.Split(uri)
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
func Split(uri string) (host string, id string, query string, fragment string, err error) {
	const scheme string = xrpcuri_internal.SchemeUnencrypted
	return xrpcuri_internal.Split(uri, scheme)
}
