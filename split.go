package xrpcuri

import (
	gourl "net/url"

	"github.com/reiver/go-erorr"
)

// Split returns the 'host', 'id', 'query', and 'fragment' of at XRPC-URI.
//
// A 'id' should be an NSID (Namespaced Identifier).
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
func Split(uri string) (host string, id string, query string, fragment string, err error) {
	const scheme string = Scheme
	return split(uri, scheme)
}

func split(uri string, scheme string) (host string, id string, query string, fragment string, err error) {
	if "" == uri {
		err = errEmptyURI
		return
	}

	var urloc *gourl.URL
	{
		urloc, err = gourl.Parse(uri)
		if nil != err {
			if e, casted := err.(*gourl.Error); casted {
				if ":" == e.URL {
					err = erorr.Errorf("xrpc: expected scheme to be %q but actually was \"\"", scheme)
				}
			}
			return
		}
		if nil == urloc {
			err = errNilURL
			return
		}
	}

	switch urloc.Scheme {
	case scheme:
		// nothing here
	default:
		err = erorr.Errorf("xrpc: expected scheme to be %q but actually was %q", scheme, urloc.Scheme)
		return
	}

	var nsid string = urloc.Path
	if 0 < len(nsid) && '/' == nsid[0] {
		nsid = nsid[1:]
	}

	host   = urloc.Host
	id = nsid
	query = urloc.RawQuery
	fragment = urloc.Fragment

	return
}


