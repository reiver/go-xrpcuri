package xrpcuri

import (
	gourl "net/url"

	"github.com/reiver/go-erorr"
)

// Split returns the 'scheme', 'host', 'collection', 'query', and 'fragment' of at XRPC-URI.
//
// A 'collection' should be an NSID (Namespaced Identifier).
//
// For example:
//
//	var uri string = "xrpc://public.api.bsky.app/app.bsky.actor.getProfile?actor=reiver.bsky.social"
//
//	scheme, host, collection, query, fragment, err := xrpcuri.Split(uri)
//	if nil != err {
//		return err
//	}
//
//	// scheme     == "xrpc"
//	// host       == "public.api.bsky.app"
//	// collection == "app.bsky.actor.getProfile"
//	// query      == "actor=reiver.bsky.social"
//	// fragment   == ""
//
// Split does NOT normalize the returned values.
func Split(uri string) (scheme string, host string, collection string, query string, fragment string, err error) {
	if "" == uri {
		err = errEmptyURI
		return
	}

	var urloc *gourl.URL
	{
		urloc, err = gourl.Parse(uri)
		if nil != err {
			return
		}
		if nil == urloc {
			err = errNilURL
			return
		}
	}

	switch urloc.Scheme {
	case Scheme:
		// nothing here
	case SchemeUnencrypted:
		// nothing here
	default:
		err = erorr.Errorf("xrpc: expected scheme to be %q or %q but was %q", Scheme, SchemeUnencrypted, urloc.Scheme)
		return
	}

	scheme = urloc.Scheme

	var nsid string = urloc.Path
	if 0 < len(nsid) && '/' == nsid[0] {
		nsid = nsid[1:]
	}

	host   = urloc.Host
	collection = nsid
	query = urloc.RawQuery
	fragment = urloc.Fragment

	return
}
