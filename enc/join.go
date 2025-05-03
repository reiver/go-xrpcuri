package xrpcurienc

import (
	"github.com/reiver/go-xrpcuri/internal"
)

func Join(authority string, id string, query string, fragment string) string {
	const scheme string = xrpcuri_internal.Scheme
	return xrpcuri_internal.Join(scheme, authority, id, query, fragment)
}
