package xrpcuri

import (
	gourl "net/url"
	"strings"

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

func split(uri string, scheme string) (authority string, id string, query string, fragment string, err error) {
	if "" == uri {
		err = errEmptyURI
		return
	}

	var kind string = "XRPC-URI"
	if SchemeUnencrypted == scheme {
		kind = "XRPC-unencrypted-URI"
	}

	var str string = uri

	{
		var prefix string = scheme + ":"

		var lenprefix int = len(prefix)
		var lenstr int = len(str)
		if lenstr < lenprefix {
			err = erorr.Errorf("xrpcuri: URI %q is not an %s because it does not begin with %q", uri, kind, prefix)
			return
		}

		var beginning string = uri[:lenprefix]

		if strings.ToLower(beginning) != prefix {
			err = erorr.Errorf("xrpcuri: URI %q is not an %s because it does not begin with %q", uri, kind, prefix)
			return
		}

		str = str[lenprefix:]
	}

	// I.e., if the URI was:
	//
	// • "at:"
	// • "aT:"
	// • "At:"
	// • "AT:"
	if "" == str {
		return
	}

	{
		const prefix string = "//"

		var lenprefix int = len(prefix)
		var lenstr int = len(str)
		if lenstr < lenprefix {
			err = erorr.Errorf("xrpcuri: %s %q is not valid because it does not have %q after \"at:\" — too short", kind, uri, prefix)
			return
		}

		var beginning string = str[:lenprefix]

		if beginning != prefix {
			err = erorr.Errorf("xrpcuri: %s %q is not valid because it does not have %q after \"at:\"", kind, uri, prefix)
			return
		}

		str = str[lenprefix:]
	}

	// authority
	{
		var index int = strings.Index(str, "/")
		if index < 0 {
			index = strings.Index(str, "?")
			if index < 0 {
				index = strings.Index(str, "#")
			}
		}
		switch {
		case index < 0:
			authority = str
			str = ""
		default:
			authority = str[:index]
			str = str[index:]
		}

		{
			unescaped, e := gourl.QueryUnescape(authority)
			if nil != e {
				err = erorr.Errorf("xrpcuri: problem hex-decoding URI %q: %w", uri, e)
				return
			}
			authority = unescaped
		}
	}

	switch str {
	case "", "/", "?", "#":
		return
	}

	// id
	{
		const prefix string = "/"

		if strings.HasPrefix(str, prefix)  {
			str = str[len(prefix):]

			var index int = strings.Index(str, "/")
			if index < 0 {
				index = strings.Index(str, "?")
				if index < 0 {
					index = strings.Index(str, "#")
				}
			}

			switch {
			case index < 0:
				id = str
				str = ""
			default:
				id = str[:index]
				str = str[index:]
			}
		}
	}

	switch str {
	case "", "/", "?", "#":
		return
	}

	switch str {
	case "", "?", "#":
		return
	}

	// query
	{
		const prefix string = "?"

		if strings.HasPrefix(str, prefix)  {
			str = str[len(prefix):]

			var index int = strings.Index(str, "#")

			switch {
			case index < 0:
				query = str
				str = ""
			default:
				query = str[:index]
				str = str[index:]
			}
		}
	}

	switch str {
	case "", "#":
		return
	}

	// fragment
	{
		const prefix string = "#"

		if strings.HasPrefix(str, prefix)  {
			str = str[len(prefix):]

			fragment = str
		}
	}

	return
}
