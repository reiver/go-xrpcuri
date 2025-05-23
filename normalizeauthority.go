package xrpcuri

import (
	"github.com/reiver/go-xrpcuri/internal"
)

// NormalizeAuthority returns the normalized form of an XRPC-URI 'authority'.
//
// The XRPC-URI 'authority' (such as "example.com") is part of an XRPC-URI (such as "xrpc://example.com").
//
// In simple language, you can think of an XRPC-URI 'authority' as being either an Internet domain-name (ex: "example.com").
//
// An example of a non-normalized XRPC-URI 'authority' would be "Example.COM".
// Normalizing that non-normalized XRPC-URI 'authority' would result in "example.com".
//
// NormalizeAuthority will leave any potential 'userinfo' in the 'authority' as is.
//
// Note that if you want to normalize a whole XRPC-URI rather than just an 'authority', then instead use [Normalize].
func NormalizeAuthority(authority string) string {
	return xrpcuri_internal.NormalizeAuthority(authority)
}
