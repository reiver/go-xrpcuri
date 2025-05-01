package xrpcuri

import (
	gourl "net/url"
)

func Join(host string, collection string, query string, fragment string) string {

	const scheme string = Scheme

	return join(scheme, host, collection, query, fragment)
}

func join(scheme string, host string, collection string, query string, fragment string) string {
	var url = gourl.URL{
		Scheme:   scheme,
		Host:     host,
		Path:     collection,
		RawQuery: query,
		Fragment: fragment,
	}

	return url.String()
}
