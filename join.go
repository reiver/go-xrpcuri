package xrpcuri

import (
	gourl "net/url"
)

func Join(host string, id string, query string, fragment string) string {

	const scheme string = Scheme

	return join(scheme, host, id, query, fragment)
}

func join(scheme string, host string, id string, query string, fragment string) string {
	var url = gourl.URL{
		Scheme:   scheme,
		Host:     host,
		Path:     id,
		RawQuery: query,
		Fragment: fragment,
	}

	return url.String()
}
