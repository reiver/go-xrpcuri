package xrpcuri

import (
	"github.com/reiver/go-xrpcuri/enc"
)

func Join(authority string, id string, query string, fragment string) string {
	return xrpcurienc.Join(authority, id, query, fragment)
}
