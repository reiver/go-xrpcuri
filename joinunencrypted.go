package xrpcuri

import (
	"github.com/reiver/go-xrpcuri/pln"
)

func JoinUnencrypted(authority string, id string, query string, fragment string) string {
	return xrpcuripln.Join(authority, id, query, fragment)
}
