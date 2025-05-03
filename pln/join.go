package xrpcuripln

import (
	"github.com/reiver/go-xrpcuri/internal"
)

func Join(host string, id string, query string, fragment string) string {
	return xrpcuri_internal.Join(Scheme, host, id, query, fragment)
}
