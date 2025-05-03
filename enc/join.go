package xrpcurienc

import (
	"github.com/reiver/go-xrpcuri/internal"
)

func Join(authority string, id string, query string, fragment string) string {
	return xrpcuri_internal.Join(Scheme, authority, id, query, fragment)
}
