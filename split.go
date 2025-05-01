package xrpcuri

import (
	gourl "net/url"

	"github.com/reiver/go-erorr"
)

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
