package xrpcuri

// SplitUnencrypted returns the 'host', 'collection', 'query', and 'fragment' of at XRPC-URI.
//
// A 'collection' should be an NSID (Namespaced Identifier).
//
// For example:
//
//	var uri string = "xrpc-unencrypted://public.api.bsky.app/app.bsky.actor.getProfile?actor=reiver.bsky.social"
//
//	host, collection, query, fragment, err := xrpcuri.SplitUnencrypted(uri)
//	if nil != err {
//		return err
//	}
//
//	// host       == "public.api.bsky.app"
//	// collection == "app.bsky.actor.getProfile"
//	// query      == "actor=reiver.bsky.social"
//	// fragment   == ""
//
// SplitUnencrypted does NOT normalize the returned values.
func SplitUnencrypted(uri string) (host string, collection string, query string, fragment string, err error) {
	const scheme string = SchemeUnencrypted
	return split(uri, scheme)
}
