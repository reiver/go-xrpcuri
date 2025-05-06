package xrpcuri

import (
	"github.com/reiver/go-xrpcuri/internal"
)

// ResolveID resolves the 'id' (which is expected to be an NSID) into an HTTP path.
//
// If `id` is "" then it returns "".
func ResolveID(id string) string {
	return xrpcuri_internal.ResolveID(id)
}
