package xrpcuri

import (
	"github.com/reiver/go-xrpcuri/internal"
)

func JoinUnencrypted(host string, id string, query string, fragment string) string {

	const scheme string = xrpcuri_internal.SchemeUnencrypted

	return join(scheme, host, id, query, fragment)
}
